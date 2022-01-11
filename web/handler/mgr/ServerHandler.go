package mgr

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/web/ws"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
)

func GetServers(c *gin.Context) {
	servers, err := data.GetServers()

	if err != nil {
		boot.Resp(c, "100100", "请求异常", err.Error())
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
	if servers == nil {
		servers = []table.Server{}
	}

	boot.Resp(c, "100200", "成功", servers)
}

func GetServer(c *gin.Context) {
	ServerName := c.Query("ServerName")

	server, err := data.GetServer(ServerName)
	if err != nil {
		boot.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	boot.Resp(c, "100200", "成功", server)
}

func DeleteServer(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	ServerName := json["ServerName"].(string)

	err := data.DeleteServer(ServerName)
	if err != nil {
		boot.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	boot.Resp(c, "100200", "成功", "")
}

func GetServerNames(c *gin.Context) {
	res, err := data.GetServersName()

	if err != nil {
		boot.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	boot.Resp(c, "100200", "成功", res)
}
