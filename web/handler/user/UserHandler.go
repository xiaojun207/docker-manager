package user

import (
	"docker-manager/data"
	"docker-manager/model"
	"docker-manager/service"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

// 用户列表
func UserListHandler(c *gin.Context, page *model.Page) {
	users, _ := service.FindUsers(page)

	boot.Resp(c, "100200", "成功", gin.H{
		"list": users,
		"page": page,
	})
}

// 用户信息
func UserInfoHandler(c *gin.Context) {
	uid := c.GetInt("uid")
	user, err := service.FindUser(uid)
	if err != nil {
		log.Println("UserInfoHandler.err:", err)
		boot.Resp(c, "103100", "获取用户信息错误", "")
		return
	}
	info := map[string]interface{}{
		"roles":        []string{user.Role},
		"introduction": user.Summary,
		"avatar":       "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png",
		"name":         user.Username,
	}
	boot.Resp(c, "100200", "成功", info)
}

type ReqUser struct {
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func AddUserHandler(c *gin.Context, req ReqUser) {
	log.Println("AddUserHandler:", req)
	err := service.AddUser(req.Nickname, req.Username, req.Mobile, req.Email, req.Password, req.Role)
	if err != nil {
		log.Println(err)
		boot.Resp(c, "100100", err.Error(), "")
		return
	}
	boot.Resp(c, "100200", "成功", "")
}

func DeleteUserHandler(c *gin.Context, req struct {
	Id int `json:"id,string"`
}) {
	uid := req.Id
	err := data.DeleteUser(uid)
	if err != nil {
		log.Println(err)
		boot.Resp(c, "100100", err.Error(), "")
		return
	}
	boot.Resp(c, "100200", "成功", "")
}

func ChangeStatusHandler(c *gin.Context, req struct {
	Username string `json:"username"`
	Status   int    `json:"status"`
}) {
	err := service.ChangeStatus(req.Username, req.Status)
	if err != nil {
		log.Println(err)
		boot.Resp(c, "100100", err.Error(), "")
		return
	}
	boot.Resp(c, "100200", "成功", "")
}
