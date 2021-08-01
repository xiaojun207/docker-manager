package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/web/ws"
	"github.com/go-basic/uuid"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
	"time"
)

func SaveAndSendTask(serverName, ch string, param map[string]interface{}) error {
	err := ws.AgentWsConnectGroup.Push(serverName, ch, param)
	if err != nil {
		return err
	}

	data.AddTask(table.Task{
		TaskId:     param["taskId"].(string),
		Name:       "task-" + ch,
		Ch:         ch,
		Code:       "000000",
		Msg:        "",
		ServerName: serverName,
		Ts:         time.Now().Unix(),
		Param:      utils2.MapToJson(param),
		Resp:       "",
		CreateDate: time.Time{},
		UpdateDate: time.Time{},
	})
	return err
}

func SendToAllServer(ch string, param map[string]interface{}) {
	reslist, err := data.GetServers()
	if err != nil {
		log.Println("SendToAllServer.err:", err)
	}
	for _, ServerName := range reslist {
		param["taskId"] = uuid.New()
		param["ServerName"] = ServerName

		if !ws.AgentConnected(ServerName.Name) {
			log.Println("ServerName:" + ServerName.Name + "服务器已离线")
			continue
		}
		err := SaveAndSendTask(ServerName.Name, ch, param)
		if err != nil {
			log.Println("err:", err)
		}
	}
}
