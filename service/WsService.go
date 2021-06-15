package service

import (
	"docker-manager/data"
	"docker-manager/web/ws"
	"time"
)

func SaveAndSendTask(serverName, ch string, param map[string]interface{}) error {
	err := ws.Push(serverName, ch, param)
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
