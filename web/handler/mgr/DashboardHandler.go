package mgr

import (
	"docker-manager/data"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
)

func GetDashboardSize(c *gin.Context) {
	containerSizeMap := data.GetContainerSize()

	d := map[string]int64{
		"task":         data.GetTaskSize(),
		"server":       data.GetServersSize(),
		"container":    containerSizeMap["totalSize"],
		"containerRun": containerSizeMap["runningSize"],
		"image":        data.GetImageSize(),
		"app":          data.GetServiceSize(),
		"follow":       containerSizeMap["followSize"],
	}
	resp.Resp(c, "100200", "成功", d)
}

func GetTaskSize(c *gin.Context) {
	resp.Resp(c, "100200", "成功", data.GetTaskSize())
}

func GetAppSize(c *gin.Context) {
	resp.Resp(c, "100200", "成功", data.GetServiceSize())
}

func GetImageSize(c *gin.Context) {
	resp.Resp(c, "100200", "成功", data.GetImageSize())
}

func GetServerSize(c *gin.Context) {
	resp.Resp(c, "100200", "成功", data.GetServersSize())
}

func GetContainerSize(c *gin.Context) {
	containerSizeMap := data.GetContainerSize()
	resp.Resp(c, "100200", "成功", containerSizeMap["totalSize"])
}
