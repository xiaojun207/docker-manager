package agent

import (
	"docker-manager/pkg/data"
	"docker-manager/pkg/data/table"
	"docker-manager/pkg/service"
	"docker-manager/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func ContainerStatsHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	AppId := c.GetHeader("AppId")
	log.Printf("%v", &json)
	ContainerId := json["ContainerId"].(string)
	service.UpdateStats(AppId, json)
	boot.Resp(c, "100200", "成功", gin.H{
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
		e.CpuStats = v["cpu_stats"].(map[string]interface{})
		e.MemoryStats = v["memory_stats"].(map[string]interface{})
		data.AddContainerStats(e)
	}
	boot.Resp(c, "100200", "成功", gin.H{})
}
