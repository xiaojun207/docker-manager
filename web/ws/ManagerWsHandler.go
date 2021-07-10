package ws

import (
	"docker-manager/web/ws/baseWs"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var ManagerWsConnectGroup = baseWs.NewWsConnectionGroup()

func WSManagerHandler(c *gin.Context) {
	containerId := c.Query("containerId")
	log.Println("WSManagerHandler.coming", ",containerId:", containerId)
	baseWs.WsHandler(c.Writer, c.Request, containerId, &ManagerWsConnectGroup, managerWsMsgHandler)
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
	default:
		break
	}
	return nil
}
