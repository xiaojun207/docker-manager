package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/utils"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"strings"
)

/**
  {
    "Containers": -1,
    "Created": 1639360189,
    "Id": "sha256:eefd0e24aa83a5bee0329a8a0260cc8cbf37781e4b308bded8616f33c3bde037",
    "Labels": null,
    "ParentId": "",
    "RepoDigests": [
      "xiaojun207/docker-agent@sha256:0a39fed07517ac0cc473d9d164430e6a10c88376a48475628391bd0e9261118b"
    ],
    "RepoTags": [
      "xiaojun207/docker-agent:1.2.1",
      "xiaojun207/docker-agent:latest"
    ],
    "SharedSize": -1,
    "Size": 13994247,
    "VirtualSize": 13994247
  },
*/
// UpdateImages
func UpdateImages(AppId string, json map[string]interface{}) {
	Name := json["Name"].(string)

	for _, tmp := range json["Images"].([]interface{}) {
		v := tmp.(map[string]interface{})
		v["AppId"] = AppId
		v["ServerName"] = Name
		var img table.Image
		utils2.MapToStruct(v, &img)
		img.AppId = AppId
		img.ImageId = v["Id"].(string)
		img.ServerName = Name
		if v["RepoDigests"] != nil {
			img.RepoDigests = strings.Join(utils.ArrInterfaceToStr(v["RepoDigests"].([]interface{})), ",")
		}
		if v["RepoTags"] != nil {
			img.RepoTags = strings.Join(utils.ArrInterfaceToStr(v["RepoTags"].([]interface{})), ",")
		}
		img.Created = int64(v["Created"].(float64))
		img.Size = int64(v["Size"].(float64))
		img.Summary = v
		data.AddImage(img)
	}
}
