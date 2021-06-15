package utils

import (
	"log"
	"testing"
)

func TestToken(t *testing.T) {
	token := TokenHelper.CreateToken("1", 1000)
	log.Println("token:", token)

	isValid, data, err := TokenHelper.ValidToken(token)
	log.Println("isValid, data, err:", isValid, data, err)
}
