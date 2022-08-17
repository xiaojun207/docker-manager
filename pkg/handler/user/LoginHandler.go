package user

import (
	"docker-manager/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func LoginHandler(c *gin.Context, req struct {
	Username string `json:"username"`
	Password string `json:"password"`
}) string {
	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		log.Println(err)
		panic(boot.NewWebError("100100", err.Error()))
	}
	return token
}

func LogoutHandler(c *gin.Context) {
	service.Logout(c)
}
