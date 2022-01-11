package user

import (
	"docker-manager/service"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func LoginHandler(c *gin.Context) string {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	username := json["username"].(string)
	password := json["password"].(string)

	//log.Println("json:", json)
	token, err := service.Login(username, password)
	if err != nil {
		log.Println(err)
		panic(boot.NewWebError("100100", err.Error()))
	}
	return token
}

func LogoutHandler(c *gin.Context) {
	service.Logout(c)
	log.Println("LogoutHandler.Success")
}
