package service

import (
	"docker-manager/data"
	"docker-manager/web/ws"
	"github.com/go-basic/uuid"
	"log"
	"time"
)

func SaveAndSendTask(serverName, ch string, param map[string]interface{}) error {
	err := ws.AgentWsConnectGroup.Push(serverName, ch, param)
	if err != nil {
		return err
	}
	taskMap := map[string]interface{}{
		"taskId":     param["taskId"],
		"ch":         ch,
		"serverName": serverName,
		"ts":         time.Now().Unix(),
		"param":      param,
		"code":       "000000",
	}
	data.Task.Store(param["taskId"].(string), taskMap)
	return err
}

func SendToAllServer(ch string, param map[string]interface{}) {
	for _, ServerName := range data.Servers.Keys() {
		param["taskId"] = uuid.New()
		param["ServerName"] = ServerName

		if !ws.AgentConnected(ServerName) {
			log.Println("ServerName:" + ServerName + "服务器已离线")
			continue
		}
		err := SaveAndSendTask(ServerName, ch, param)
		if err != nil {
			log.Println("err:", err)
		}
	}
}
