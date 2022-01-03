package agent

import (
	"docker-manager/service"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
)

func RegDockerHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	json["PrivateIp"] = c.GetHeader("PrivateIp")
	json["HostIp"] = c.GetHeader("HostIp")
	json["PublicIp"] = utils.GetRemoteIP(c)

	err := service.DockerReg(json)

	Name := json["Name"].(string)
	id := json["ID"].(string)
	//AppId := c.GetHeader("AppId")
	log.Println("reg server:", id, " Name:", Name)
	if err != nil {
		log.Println("add server.err:", err, " Name:", Name)
	}

	resp.Resp(c, "100200", "成功", gin.H{
		"name": Name,
		"id":   id,
	})
}
