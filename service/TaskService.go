package service

import (
	"docker-manager/data"
	"docker-manager/dto"
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func PublishAppTask(appInfo dto.AppInfo) error {
	log.Println("PublishAppTask.appInfo:", utils.StructToJson(appInfo))
	data.AddAppInfo(appInfo.Name, appInfo)

	for _, serverName := range appInfo.ServerNames {
		param := map[string]interface{}{
			"taskId":        uuid.New(),
			"serverName":    serverName,
			"imageName":     appInfo.Image,
			"containerName": appInfo.Name,
			"ports":         appInfo.Ports,
			"volumes":       appInfo.Volumes,
			"env":           appInfo.Env, // {"appversion=1.0.1"}
			"memory":        appInfo.Memory,
			"logType":       appInfo.LogType,
			"logConfig":     appInfo.LogConfig,
		}
		ch := "docker.container.run"

		err := SaveAndSendTask(serverName, ch, param)
		log.Println("SaveAndSendTask.err:", err)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
