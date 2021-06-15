package user

import (
	"docker-manager/service"
	"docker-manager/web/handler/agent"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
)

/**
eg.2
*/
func LoginHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	username := json["username"].(string)
	password := json["password"].(string)

	token, err := service.Login(username, password)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", token)
}

func LogoutHandler(c *gin.Context) {
	resp.Resp(c, "100200", "成功", "")
}

func UserInfoHandler(c *gin.Context) {
	log.Println("Version Handler:", agent.Version)
	info := map[string]interface{}{
		"roles":        []string{"admin"},
		"introduction": "I am a super administrator",
		"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"name":         "Super Admin",
	}

	resp.Resp(c, "100200", "成功", info)
}

/**
eg.3
*/
func VersionHandler(c *gin.Context) {
	log.Println("Version Handler:", agent.Version)
	resp.Resp(c, "100200", "成功", "Version:"+agent.Version)
}
