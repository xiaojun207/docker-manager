package service

import (
	"encoding/json"
	"fmt"
	"github.com/xiaojun207/go-base-utils/utils"
	"strings"
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
	lastVersion := "0.0.0"
	for _, m := range d {
		v := m["name"].(string)
		if v == "latest" {
			continue
		}
		if VersionCompare(lastVersion, v) < 0 {
			lastVersion = v
		}
	}
	return lastVersion
}

// eg.: 1.0.1
// return： -1 0 1
func VersionCompare(ver1, ver2 string) int {
	arrVer1 := strings.Split(ver1, ".")
	arrVer2 := strings.Split(ver2, ".")
	len1, len2 := len(arrVer1), len(arrVer2)
	minLen := utils.If(len1 < len2, len1, len2).(int)
	for i := 0; i < minLen; i++ {
		v1 := utils.StrToInt(arrVer1[i])
		v2 := utils.StrToInt(arrVer2[i])
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	if len1 > len2 {
		return 1
	} else if len1 < len2 {
		return -1
	} else {
		return 0
	}
}
