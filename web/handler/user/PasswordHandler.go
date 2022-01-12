package user

import (
	"docker-manager/service"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func AlterPasswordHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	uid := c.GetInt("uid")
	oldPassword := json["OldPassword"].(string)
	newPassword := json["NewPassword"].(string)

	err := service.AlterPassword(uid, oldPassword, newPassword)
	if err != nil {
		log.Println(err)
		boot.Resp(c, "100100", err.Error(), "")
		return
	}
	boot.Resp(c, "100200", "成功", "")
}

func ForgetPasswordHandler(c *gin.Context, req struct {
	Username string `json:"username"`
	Code     string `json:"code" binding:"required"`
	CodeType string `json:"codeType"`
}) {
	user, err := service.FindUserByUsername(req.Username)
	if err != nil {
		boot.Resp(c, "100100", "验证码错误", "")
		return
	}
	uid := user.Id
	if !service.CheckCode(uid, req.Code) {
		boot.Resp(c, "100100", "验证码错误", "")
		return
	}

	newPassword, err := service.ResetPassword(req.Username)
	if err != nil {
		log.Println(err)
		boot.Resp(c, "100100", err.Error(), "")
		return
	}
	boot.Resp(c, "100200", "成功", newPassword)
}

func ResetPasswordHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	username := json["username"].(string)

	newPassword, err := service.ResetPassword(username)
	if err != nil {
		log.Println(err)
		boot.Resp(c, "100100", err.Error(), "")
		return
	}
	boot.Resp(c, "100200", "成功", newPassword)
}
