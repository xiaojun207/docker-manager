package agent

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
)

func RegDockerHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	Name := json["Name"].(string)
	id := json["ID"].(string)
	//AppId := c.GetHeader("AppId")

	server := table.Server{
		Name:            Name,
		OSType:          json["OSType"].(string),
		OperatingSystem: json["OperatingSystem"].(string),
		KernelVersion:   json["KernelVersion"].(string),
		DockerVersion:   json["ServerVersion"].(string),
		Running:         int(json["ContainersRunning"].(float64)),
		Containers:      int(json["Containers"].(float64)),
		Cpu:             int(json["NCPU"].(float64)),
		Memory:          int64(json["MemTotal"].(float64)),
		Images:          int(json["Images"].(float64)),
		//PrivateIp:  "",
		//PublicIp:   "",
		//State:  "",
		Summary: json,
	}
	err := data.AddServer(server)
	log.Println("reg server:", id, " Name:", Name)
	if err != nil {
		log.Println("add server.err:", err, " Name:", Name)
	}

	resp.Resp(c, "100200", "成功", gin.H{
		"name": Name,
		"id":   id,
	})
}

func ContainersHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	AppId := c.GetHeader("AppId")

	id := json["ID"].(string)

	service.UpdateServerContainer(AppId, json)

	resp.Resp(c, "100200", "成功", gin.H{
		"id": id,
	})
}
