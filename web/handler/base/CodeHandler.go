package base

import (
	"docker-manager/service"
	"docker-manager/web/resp"
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
)

func SendCodeHandler(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	log.Println("SendCodeHandler:", json)
	username := json["username"].(string)
	codeType := json["codeType"].(string)
	key := json["key"].(string)

	user, err := service.FindUserByUsername(username)
	if err != nil {
		resp.Resp(c, "100100", "用户名错误", "")
		return
	}
	code := service.SendCode(user.Id, codeType, key)

	log.Println("SendCode, username:", ",code:", code)
	resp.Resp(c, "100200", "成功", "")
}
