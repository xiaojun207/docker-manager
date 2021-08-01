package agent

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func ContainerStatsHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	//AppId := c.GetHeader("AppId")
	log.Printf("%v", &json)
	ContainerId := json["ContainerId"].(string)
	Name := json["Name"].(string)
	Time := json["Time"].(string)

	for _, t := range json["Stats"].([]interface{}) {
		v := t.(map[string]interface{})
		v["ContainerId"] = v["id"].(string)
		v["ServerName"] = Name
		v["Time"] = Time
		v["name"] = utils.TrimContainerName(v["name"].(string))
		var stats table.ContainerStats
		utils2.MapToStruct(v, &stats)
		data.AddContainerStats(stats)
	}
	resp.Resp(c, "100200", "成功", gin.H{
		"id": ContainerId,
	})
}

func ContainersStatsHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	//AppId := c.GetHeader("AppId")
	//log.Printf("%v", utils2.MapToJson(json))
	//ID := json["ID"].(string)
	serverName := json["Name"].(string)
	Time := json["Time"].(float64)
	stats := json["Stats"].([]interface{})

	for _, t := range stats {
		v := t.(map[string]interface{})

		//log.Printf("%v", utils2.MapToJson(v))
		v["ContainerId"] = v["id"].(string)
		v["name"] = utils.TrimContainerName(v["name"].(string))
		v["ServerName"] = serverName
		v["Time"] = Time

		var e table.ContainerStats
		utils2.MapToStruct(v, &e)
		data.AddContainerStats(e)
	}
	resp.Resp(c, "100200", "成功", gin.H{})
}
