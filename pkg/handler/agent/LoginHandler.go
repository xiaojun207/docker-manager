package agent

import (
	"docker-manager/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

/**
eg.2
*/
func LoginHandler(c *gin.Context, req struct {
	Username string `json:"username"`
	Password string `json:"password"`
}) {
	token, err := service.LoginAgent(req.Username, req.Password)
	if err != nil {
		log.Println(err)
		boot.Resp(c, "100100", err.Error(), "")
		return
	}
	boot.Resp(c, "100200", "成功", token)
}
