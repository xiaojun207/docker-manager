package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/web/ws"
	"docker-manager/web/ws/baseWs"
	"github.com/go-basic/uuid"
	"log"
	"time"
)

func init() {
	ws.AgentWsConnectGroup.MsgHandler = agentMsgHandler
	ws.ManagerWsConnectGroup.MsgHandler = managerWsMsgHandler

	ws.AgentWsConnectGroup.OnConnected = func(id interface{}) {
		data.UpdateServerState(id.(string), "connected")
	}

	ws.AgentWsConnectGroup.OnDisconnected = func(id interface{}) {
		data.UpdateServerState(id.(string), "disconnect")
	}

}

func managerWsMsgHandler(msg *baseWs.WsMsg, conn *baseWs.Connection) error {
	tmp := msg.Data
	d, ok := tmp.(map[string]interface{})
	if !ok {
	}
	log.Println("ManagerWsMsgHandler.d:", d)
	switch msg.Channel {
	case baseWs.CH_PING:
		return conn.Pong()
	case baseWs.CH_PONG:
		conn.LastPongTime = time.Now().UnixNano() / 1e6
		break
	case "docker.container.stats":
		containerId := d["cId"].(string)
		res := map[string]interface{}{
			"ts":   d["ts"],
			"line": d["line"],
		}
		err := ws.ManagerWsConnectGroup.Push(containerId, "docker.container.stats", res)
		if err != nil {

		}
		break
	default:
		break
	}
	return nil
}

func agentMsgHandler(msg *baseWs.WsMsg, conn *baseWs.Connection) error {
	conn.LastDataTime = time.Now().UnixNano() / 1e6
	tmp := msg.Data
	AppId := conn.Headers["AppId"]
	d, ok := tmp.(map[string]interface{})
	if !ok {
	}
	switch msg.Channel {
	case baseWs.CH_PING:
		return conn.Pong()
	case baseWs.CH_PONG:
		conn.LastPongTime = conn.LastDataTime
		break
	case "docker.info.systemTime":

		break
	case "docker.info":

		break
	case "docker.task.ack":
		// {"code":code, "msg": err, "taskId":taskId, "resp": resp }
		log.Println("docker.task.ack :", d)
		//d := utils.MapInterfaceToString(d)
		taskId := d["taskId"].(string)
		code := d["code"].(string)
		data.UpdateTask(taskId, code, d["msg"].(string), d["resp"].(map[string]interface{}))
		break
	case "docker.container.stats":
		UpdateStats(AppId, d)
		break
	case "docker.container.log.line":
		// {"code":code, "msg": err, "taskId":taskId, "resp": resp }
		containerId := d["cId"].(string)
		res := map[string]interface{}{
			"ts":   d["ts"],
			"line": d["line"],
		}
		err := ws.ManagerWsConnectGroup.Push(containerId, "docker.container.log.line", res)
		if err != nil {
			//param := map[string]interface{}{
			//	"taskId":      uuid.New(),
			//	"containerId": containerId,
			//}
			//serverName := data.GetServerNameByContainerId(containerId)

			//ch := "docker.container.log.follow.close"
			//err := SaveAndSendTask(serverName, ch, param)
			//log.Println("web.log.client is close, and send cmd to agent ,serverName:", serverName, ",push.err:", err)
		}
		break
	case "docker.container.list":
		UpdateServerContainer(AppId, d)
		break
	case "docker.image.list":
		UpdateImages(AppId, d)
		break
	default:
		break
	}
	return nil
}

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
		Param:      param,
		//Resp:       "",
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
		param["ServerName"] = ServerName.Name

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
