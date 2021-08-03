package ws

import (
	"docker-manager/data"
	"docker-manager/web/ws/baseWs"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var AgentWsConnectGroup = baseWs.NewWsConnectionGroup()

func WSAgentHandler(c *gin.Context) {
	serverName := c.GetHeader("ServerName")
	log.Println("WSAgentHandler.coming", ",ServerName:", serverName)
	baseWs.WsHandler(c.Writer, c.Request, serverName, &AgentWsConnectGroup, agentMsgHandler)
}

func AgentConnected(id string) bool {
	return AgentWsConnectGroup.IsConnected(id)
}

func agentMsgHandler(msg *baseWs.WsMsg, conn *baseWs.Connection) error {
	tmp := msg.Data
	d, ok := tmp.(map[string]interface{})
	if !ok {
	}
	switch msg.Channel {
	case baseWs.CH_PING:
		return conn.Pong()
	case baseWs.CH_PONG:
		conn.LastPongTime = time.Now().UnixNano() / 1e6
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
		containerId := d["cId"].(string)
		res := map[string]interface{}{
			"ts":   d["ts"],
			"line": d["line"],
		}
		err := ManagerWsConnectGroup.Push(containerId, "docker.container.stats", res)
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
		err := ManagerWsConnectGroup.Push(containerId, "docker.container.log.line", res)
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
