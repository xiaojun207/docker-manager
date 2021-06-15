package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
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
	webName string
	secret  []byte
}

var TokenHelper = TokenFactory{
	webName: "GoApp",
	secret:  []byte("50c9b0ac9f38836d4efb0dfa7504455f8304d8a16a7034f11e144e69e277a713"),
}

func CreateToken(data string) string {
	return TokenHelper.CreateToken(data, 60*60*24)
}

func ValidToken(token string) (isValid bool, data string, err error) {
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
	now := time.Now().Unix()

	var tk jwt.StandardClaims

	tk.Issuer = e.webName
	tk.Audience = "web"
	tk.Subject = data
	tk.IssuedAt = now
	tk.ExpiresAt = now + expiration

	token := e.newWithClaims(tk)
	tokenString, _ := token.SignedString(e.secret)
	return tokenString
}

func (e *TokenFactory) ValidToken(tokenString string) (isValid bool, data string, err error) {
	var tk jwt.StandardClaims
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
