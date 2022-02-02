package mgr

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/model"
	"docker-manager/web/ws"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
)

// 获取宿主机列表
func GetServers(c *gin.Context, page model.Page) {
	servers, err := data.GetServers(&page)

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

	boot.Resp(c, "100200", "成功", gin.H{
		"list": servers,
		"page": page,
	})
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

func DeleteServer(c *gin.Context, req struct {
	ServerName string `json:"ServerName"`
}) {
	err := data.DeleteServer(req.ServerName)
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
