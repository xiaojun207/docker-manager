package user

import (
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
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
		resp.Resp(c, "100100", err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", "")
}

func ForgetPasswordHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	username := json["username"].(string)
	code := json["code"].(string)
	user, err := service.FindUserByUsername(username)
	if err != nil {
		resp.Resp(c, "100100", "验证码错误", "")
		return
	}
	uid := user.Id
	if !service.CheckCode(uid, code) {
		resp.Resp(c, "100100", "验证码错误", "")
		return
	}

	newPassword, err := service.ResetPassword(username)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", newPassword)
}

func ResetPasswordHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	username := json["username"].(string)

	newPassword, err := service.ResetPassword(username)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", newPassword)
}
