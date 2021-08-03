package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func PublishAppTask(serverNames []string, service table.Service) error {
	log.Println("PublishAppTask.appInfo:", utils.StructToJson(service))
	data.AddService(service)

	for _, serverName := range serverNames {
		param := map[string]interface{}{
			"taskId":        uuid.New(),
			"serverName":    serverName,
			"imageName":     service.Image,
			"containerName": service.Name,
			"ports":         service.Ports,
			"volumes":       service.Vol,
			"env":           service.Env, // {"appversion=1.0.1"}
			"memory":        service.Memory,
			//"logType":       service.LogType,
			//"logConfig":     service.LogConfig,
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
