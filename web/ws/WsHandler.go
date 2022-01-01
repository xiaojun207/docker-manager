package ws

import (
	"docker-manager/web/ws/baseWs"
	"github.com/gin-gonic/gin"
	"log"
)

var AgentWsConnectGroup = baseWs.NewWsConnectionGroup()
var ManagerWsConnectGroup = baseWs.NewWsConnectionGroup()
var ManagerExecWsConnectGroup = baseWs.NewWsConnectionGroup()

func WSAgentHandler(c *gin.Context) {
	serverName := c.GetHeader("ServerName")
	AppId := c.GetHeader("AppId")
	log.Println("WSAgentHandler.coming", ",ServerName:", serverName, ",AppId:", AppId)
	baseWs.WsHandler(c.Writer, c.Request, serverName, &AgentWsConnectGroup)
}

func WSManagerHandler(c *gin.Context) {
	containerId := c.Query("containerId")
	log.Println("WSManagerHandler.coming", ",containerId:", containerId)
	baseWs.WsHandler(c.Writer, c.Request, containerId, &ManagerWsConnectGroup)
}

func WSManagerExecHandler(c *gin.Context) {
	containerId := c.Query("containerId")
	cmd := c.Query("cmd")
	c.Request.Header.Add("containerId", containerId)
	c.Request.Header.Add("cmd", cmd)
	log.Println("WSManagerExecHandler.coming", ",containerId:", containerId, ",cmd:", cmd)
	baseWs.WsHandler(c.Writer, c.Request, containerId, &ManagerExecWsConnectGroup)
}

func AgentConnected(id string) bool {
	return AgentWsConnectGroup.IsConnected(id)
}

func AgentConnectedCount() int {
	return AgentWsConnectGroup.GetCount()
}

func AgentLastDataTime(id string) int64 {
	return AgentWsConnectGroup.LastDataTime(id)
}
