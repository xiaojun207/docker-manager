package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
)

func UpdateImages(AppId string, json map[string]interface{}) {
	//log.Printf("%v",&json)
	Name := json["Name"].(string)

	for _, v := range json["Images"].([]interface{}) {
		v.(map[string]interface{})["AppId"] = AppId
		v.(map[string]interface{})["ServerName"] = Name
		var i table.Image
		utils2.MapToStruct(v.(map[string]interface{}), &i)
		data.AddImage(i)
	}
}
