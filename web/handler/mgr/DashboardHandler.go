package mgr

import (
	"docker-manager/data"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
)

func GetDashboardSize(c *gin.Context) {
	containerRun := getContainerSize("running")
	container := getContainerSize("")
	d := map[string]int{
		"task":         data.Task.Size(),
		"server":       data.Servers.Size(),
		"container":    container,
		"containerRun": containerRun,
		"image":        data.Images.Size(),
		"app":          data.AppInfos.Size(),
		"follow":       data.GetFollowSize(),
	}
	resp.Resp(c, "100200", "成功", d)
}

func GetTaskSize(c *gin.Context) {
	resp.Resp(c, "100200", "成功", data.Task.Size())
}

func GetAppSize(c *gin.Context) {
	resp.Resp(c, "100200", "成功", data.AppInfos.Size())
}

func GetImageSize(c *gin.Context) {
	resp.Resp(c, "100200", "成功", data.Images.Size())
}

func GetServerSize(c *gin.Context) {
	resp.Resp(c, "100200", "成功", data.Servers.Size())
}

func getContainerSize(state string) int {
	res := 0
	data.Containers.ForEach(func(_ string, arr []byte) {
		tmps := []map[string]interface{}{}
		utils2.JsonBytesToStruct(arr, &tmps)
		if state == "" {
			res = res + len(tmps)
			return
		}
		for _, v := range tmps {
			container := v
			c_state := container["State"].(string)

			if c_state != state {
				continue
			}

			res = res + 1
		}
	})
	return res
}

func GetContainerSize(c *gin.Context) {
	resp.Resp(c, "100200", "成功", getContainerSize(""))
}
