package data

import (
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
	"testing"
)

func TestCreatePassword(t *testing.T) {
	slat := utils.Sha256(uuid.New())
	password := "a8ceec2ef7294f46f6fd4fea32c0aa4d6d76dd27c86b4f7c470ee5b568382057"
	dbPassword := utils.Sha256(password + slat)
	log.Println("slat:", slat)
	log.Println("password:", password)
	log.Println("dbPassword:", dbPassword)
}

func TestCreatePassword2(t *testing.T) {
	slat := utils.Sha256(uuid.New())
	slat = "3367db0083b430c6986a538181f5462a9c8f78af9658067605e58496f475bf0f"
	password := "72d60f075a2341ab"
	dbPassword := utils.Sha256(password + slat)
	log.Println("slat:", slat)
	log.Println("password:", password)
	log.Println("dbPassword:", dbPassword)
}
