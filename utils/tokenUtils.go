package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
	"sync"
	"time"
)

var (
	UseCacheMap = true
	// 一些常量
	ErrorTokenExpired     error = boot.NewWebError("105101", "Token is expired")
	ErrorTokenNotValidYet error = boot.NewWebError("105101", "Token not active yet")
	ErrorTokenMalformed   error = boot.NewWebError("105101", "That's not even a token")
	ErrorTokenInvalid     error = boot.NewWebError("105101", "Couldn't handle this token:")
	ErrorTokenNotExists   error = boot.NewWebError("105101", "Token not exists")
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
	if UseCacheMap {
		_, ok := TokenHelper.cacheMap.Load(token)
		if !ok {
			return false, "", ErrorTokenNotExists
		}
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
				err = ErrorTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				err = ErrorTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				err = ErrorTokenNotValidYet
			} else {
				err = ErrorTokenInvalid
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
