package ws

import (
	"docker-manager/data"
	"log"
	"time"
)

func WsMsgHandler(msg *WsMsg, conn *Connection) error {
	tmp := msg.Data
	d, ok := tmp.(map[string]interface{})
	if !ok {
	}
	switch msg.Channel {
	case CH_PING:
		return conn.Pong()
	case CH_PONG:
		conn.lastPongTime = time.Now().UnixNano() / 1e6
		break
	case "docker.task.ack":
		// {"code":code, "msg": err, "taskId":taskId, "resp": resp }
		log.Println("docker.task.ack :", d)
		//d := utils.MapInterfaceToString(d)
		taskId := d["taskId"].(string)
		code := d["code"].(string)
		task := data.Task.LoadMap(taskId)
		if task != nil {
			task["code"] = code
			task["msg"] = d["msg"]
			task["resp"] = d["resp"]
			data.Task.Store(taskId, task)
		}
		break
	case "docker.container.stats":
		containerId := d["cId"].(string)
		res := map[string]interface{}{
			"ts":   d["ts"],
			"line": d["line"],
		}
		err := Push(containerId, "docker.container.stats", res)
		if err != nil {

		}
		break
	case "docker.container.log.line":
		// {"code":code, "msg": err, "taskId":taskId, "resp": resp }
		containerId := d["cId"].(string)
		res := map[string]interface{}{
			"ts":   d["ts"],
			"line": d["line"],
		}
		err := Push(containerId, "docker.container.log.line", res)
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
	default:
		break
	}
	return nil
}

func SaveAndSendTask(serverName, ch string, param map[string]interface{}) error {
	err := Push(serverName, ch, param)
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
