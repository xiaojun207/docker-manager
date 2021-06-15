package agent

import (
	"docker-manager/data"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
)

func ContainerStatsHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	//AppId := c.GetHeader("AppId")
	//log.Printf("%v",&json)
	ContainerId := json["ContainerId"].(string)
	Name := json["Name"].(string)
	Time := json["Time"].(string)

	for _, v := range json["Stats"].([]interface{}) {
		v.(map[string]interface{})["ContainerId"] = ContainerId
		v.(map[string]interface{})["ServerName"] = Name
		v.(map[string]interface{})["Time"] = Time
	}

	data.Stats.Store(ContainerId, json["Stats"])

	resp.Resp(c, "100200", "成功", gin.H{
		"id": ContainerId,
	})
}

func ContainersStatsHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	//AppId := c.GetHeader("AppId")
	//log.Printf("%v",&json)
	//ID := json["ID"].(string)
	serverName := json["Name"].(string)
	Time := json["Time"].(float64)
	stats := json["Stats"].([]interface{})

	for _, t := range stats {
		v := t.(map[string]interface{})
		v["ServerName"] = serverName
		v["Time"] = Time
	}

	data.Stats.Store(serverName, stats)
	resp.Resp(c, "100200", "成功", gin.H{})
}
