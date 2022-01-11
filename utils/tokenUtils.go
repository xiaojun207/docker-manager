package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"sync"
	"time"
)

var (
	// 一些常量
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type TokenFactory struct {
	webName  string
	secret   []byte
	cacheMap sync.Map
}

var TokenHelper TokenFactory

func InitTokenHelper(tokenSecret string) TokenFactory {
	TokenHelper = TokenFactory{
		webName: "DockerManager",
		secret:  []byte(tokenSecret),
	}
	return TokenHelper
}

func CreateToken(data string) string {
	expiration := int64(60 * 60 * 24)
	token := TokenHelper.CreateToken(data, expiration)
	// 缓存token到内存中，如果集群模式下，请存入redis
	TokenHelper.cacheMap.Store(token, "1")
	return token
}

func RemoveToken(token string) {
	TokenHelper.cacheMap.Delete(token)
}

func ValidToken(token string) (isValid bool, data string, err error) {
	_, ok := TokenHelper.cacheMap.Load(token)
	if !ok {
		return false, "", errors.New("Token not exists")
	}
	return TokenHelper.ValidToken(token)
}

func (e *TokenFactory) newWithClaims(claims jwt.Claims) *jwt.Token {
	method := jwt.GetSigningMethod("HS512")
	return &jwt.Token{
		Header: map[string]interface{}{
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
}

/**
expiration 秒级
*/
func (e *TokenFactory) CreateToken(data string, expiration int64) string {
	now := time.Now()
	var tk jwt.RegisteredClaims

	tk.Issuer = e.webName
	tk.Audience = jwt.ClaimStrings{"web"}
	tk.Subject = data
	tk.IssuedAt = jwt.NewNumericDate(now)
	tk.ExpiresAt = jwt.NewNumericDate(now.Add(time.Second * time.Duration(expiration)))

	token := e.newWithClaims(tk)
	tokenString, _ := token.SignedString(e.secret)
	return tokenString
}

func (e *TokenFactory) ValidToken(tokenString string) (isValid bool, data string, err error) {
	var tk jwt.RegisteredClaims
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return e.secret, nil
	}
	token, err := jwt.ParseWithClaims(tokenString, &tk, keyFunc)

	if err != nil {
		log.Println("Parse token error:", err)

		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err = TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				err = TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				err = TokenNotValidYet
			} else {
				err = TokenInvalid
			}
		}
		return false, "", err
	}

	isValid = token.Valid
	if token.Valid {
		data = tk.Subject
	}
	return
}
