package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/utils"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func UpdateStats(AppId string, json map[string]interface{}) {
	//AppId := c.GetHeader("AppId")
	log.Printf("%v", &json)
	//ContainerId := json["ContainerId"].(string)
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
}
