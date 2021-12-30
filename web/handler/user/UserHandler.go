package user

import (
	"docker-manager/data"
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
)

func UserListHandler(c *gin.Context) {
	users, _ := service.FindUsers()
	resp.Resp(c, "100200", "成功", users)
}

func UserInfoHandler(c *gin.Context) {
	uid := c.GetInt("uid")
	user, err := service.FindUser(uid)
	if err != nil {
		log.Println("UserInfoHandler.err:", err)
		resp.Resp(c, "103100", "获取用户信息错误", "")
		return
	}
	info := map[string]interface{}{
		"roles":        []string{user.Role},
		"introduction": user.Summary,
		"avatar":       "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png",
		"name":         user.Username,
	}
	resp.Resp(c, "100200", "成功", info)
}

func AddUserHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	nickname := json["nickname"].(string)
	username := json["username"].(string)
	mobile := json["mobile"].(string)
	email := json["email"].(string)
	role := json["role"].(string)
	password := json["password"].(string)

	log.Println("AddUserHandler:", json)
	err := service.AddUser(nickname, username, mobile, email, password, role)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", "")
}

func DeleteUserHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	uid := int(json["id"].(float64))

	err := data.DeleteUser(uid)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", "")
}

func ChangeStatusHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	username := json["username"].(string)
	status := json["status"].(float64)

	err := service.ChangeStatus(username, int(status))
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", "")
}
