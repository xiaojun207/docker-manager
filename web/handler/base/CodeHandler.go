package base

import (
	"docker-manager/conf"
	"docker-manager/service"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

// 验证码发送
func SendCodeHandler(c *gin.Context) interface{} {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	log.Println("SendCodeHandler:", json)
	username := json["username"].(string)
	codeType := json["codeType"].(string)
	key := json["key"].(string)

	user, err := service.FindUserByUsername(username)
	if err != nil {
		return boot.NewWebError("100100", "用户名错误")
	}
	code := service.SendCode(user.Id, codeType, key)
	if utils.StrToBool(conf.ConsoleCode) {
		log.Println("SendCode, username:", username, ",code:", code)
	} else {
		log.Println("SendCode, username:", username, ",code:", "ConsoleCode is off，you can set env \"-e consoleCode=true\", or set in file of config.yml")
	}
	return nil
}
