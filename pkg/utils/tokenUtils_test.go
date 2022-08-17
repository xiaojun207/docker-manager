package utils

import (
	"github.com/go-basic/uuid"
	"log"
	"testing"
)

func TestToken(t *testing.T) {
	InitTokenHelper(uuid.New())
	token := TokenHelper.CreateToken("1", 1000)
	log.Println("token:", token)

	isValid, data, err := TokenHelper.ValidToken(token)
	log.Println("isValid, data, err:", isValid, data, err)
}
