package base

import (
	"docker-manager/pkg/conf"
	"docker-manager/pkg/service"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

// 验证码发送
func SendCodeHandler(c *gin.Context, req struct {
	Username string `json:"username"`
	CodeType string `json:"codeType"`
	Key      string `json:"key"`
}) interface{} {
	log.Println("SendCodeHandler:", req)

	user, err := service.FindUserByUsername(req.Username)
	if err != nil {
		return boot.NewWebError("100100", "用户名错误")
	}
	code := service.SendCode(user.Id, req.CodeType, req.Key)
	if utils.StrToBool(conf.ConsoleCode) {
		log.Println("SendCode, username:", req.Username, ",code:", code)
		return nil
	} else {
		log.Println("SendCode, username:", req.Username, ",code:", "ConsoleCode is off，you can set env \"-e consoleCode=true\", or set in file of config.yml")
		return boot.NewWebError("100100", "控制台验证码未开启，请配置参数consoleCode为true后重新启动")
	}
}
