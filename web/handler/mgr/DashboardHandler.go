package mgr

import (
	"docker-manager/data"
	"docker-manager/web/ws"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
)

func GetDashboardSize(c *gin.Context) {
	containerSizeMap := data.GetContainerSize()

	d := map[string]int64{
		"task":            data.GetTaskSize(),
		"server":          data.GetServersSize(),
		"serverConnected": int64(ws.AgentConnectedCount()),
		"container":       containerSizeMap["totalSize"],
		"containerRun":    containerSizeMap["runningSize"],
		"image":           data.GetImageSize(),
		"app":             data.GetServiceSize(),
		"follow":          containerSizeMap["followSize"],
	}
	boot.Resp(c, "100200", "成功", d)
}

func GetTaskSize(c *gin.Context) {
	boot.Resp(c, "100200", "成功", data.GetTaskSize())
}

func GetAppSize(c *gin.Context) {
	boot.Resp(c, "100200", "成功", data.GetServiceSize())
}

func GetImageSize(c *gin.Context) {
	boot.Resp(c, "100200", "成功", data.GetImageSize())
}

func GetServerSize(c *gin.Context) {
	boot.Resp(c, "100200", "成功", data.GetServersSize())
}

func GetContainerSize(c *gin.Context) {
	containerSizeMap := data.GetContainerSize()
	boot.Resp(c, "100200", "成功", containerSizeMap["totalSize"])
}
