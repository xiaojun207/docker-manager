package user

import (
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
)

func LoginHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	username := json["username"].(string)
	password := json["password"].(string)

	//log.Println("json:", json)
	token, err := service.Login(username, password)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", token)
}

func LogoutHandler(c *gin.Context) {
	service.Logout("")
	resp.Resp(c, "100200", "成功", "")
}
