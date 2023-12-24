package agent

import (
	"docker-manager/pkg/service"
	"docker-manager/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
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

	boot.Resp(c, boot.CodeSuccess, "成功", gin.H{
		"name": Name,
		"id":   id,
	})
}
