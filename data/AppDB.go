package data

import (
	"encoding/json"
	"log"
)

func AddAppGroup(appname, serverName string) {
	arrByte, err := AppGroup.Load(appname)
	if err != nil {
		log.Println("AddAppGroup.err:", err)
	}
	serverNames := []string{}
	json.Unmarshal(arrByte, &serverNames)
	for _, s := range serverNames {
		if s == serverName {
			return
		}
	}
	serverNames = append(serverNames, serverName)
	AppGroup.Store(appname, serverNames)
}

// 	AppInfos.Store("docker-agent", infoMap)
//	map[string]interface{}{
//		"Image": "xiaojun207/docker-agent:latest",
//		"Ports": []string{},
//		"Volumes": []string{},
//		"Mounts": []string{},
//		"env": []string{}
//	}
func AddAppInfo(appname string, info interface{}) {
	if AppInfos.Has(appname) {
		return
	}
	AppInfos.Store(appname, info)
}

func GetAppInfo(name string, info interface{}) {
	data, err := AppInfos.Load(name)
	if err != nil {
		log.Println("GetAppInfo.err:", err)
	}
	log.Println("GetAppInfo:", string(data))
	json.Unmarshal(data, info)
}
