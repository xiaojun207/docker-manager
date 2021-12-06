package service

import (
	"encoding/json"
	"fmt"
	"github.com/xiaojun207/go-base-utils/utils"
)

func GetLatestTag() string {
	image := "xiaojun207/docker-manager"
	url := "https://registry.hub.docker.com/v1/repositories/" + image + "/tags"
	rssp := utils.Get(url)
	var d []map[string]interface{}
	err := json.Unmarshal([]byte(rssp), &d)
	if err != nil {
		fmt.Println("JsonToMap err: ", err)
	}
	lastVersion := ""
	for _, m := range d {
		v := m["name"].(string)
		if v == "latest" {
			continue
		}
		if lastVersion < v {
			lastVersion = v
		}
	}
	return lastVersion
}
