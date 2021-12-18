package mgr

import (
	"docker-manager/data"
	"docker-manager/web/resp"
	"docker-manager/web/ws"
	"github.com/gin-gonic/gin"
)

func GetServers(c *gin.Context) {
	servers, err := data.GetServers()

	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	for i, server := range servers {
		if ws.AgentConnected(server.Name) {
			server.State = "connected"
		} else {
			server.State = "disconnect"
		}
		server.Summary = nil
		server.LastDataTime = ws.AgentLastDataTime(server.Name)
		servers[i] = server
	}
	resp.Resp(c, "100200", "成功", servers)
}

func GetServer(c *gin.Context) {
	ServerName := c.Query("ServerName")

	server, err := data.GetServer(ServerName)
	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	resp.Resp(c, "100200", "成功", server)
}

func GetServerNames(c *gin.Context) {
	res, err := data.GetServersName()

	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	resp.Resp(c, "100200", "成功", res)
}
