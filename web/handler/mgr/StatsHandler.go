package mgr

import (
	"docker-manager/data"
	"docker-manager/service"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetStats(c *gin.Context) {
	serverNames := c.QueryArray("serverNames[]")
	ContainerNames := c.QueryArray("ContainerNames[]")
	Follow := c.Query("Follow")
	log.Println("serverNames:", serverNames, ",ContainerNames:", ContainerNames, ",Follow:", Follow)
	res := []interface{}{}

	statss, _ := data.GetContainerStats()
	for _, stats := range statss {
		if len(serverNames) > 0 && !utils.StrInArr(serverNames, stats.ServerName) {
			continue
		}
		c_name := utils.TrimContainerName(stats.Name)
		if len(ContainerNames) > 0 && !utils.StrInArr(ContainerNames, c_name) {
			continue
		}
		if len(Follow) > 0 {
			f, _ := strconv.ParseBool(Follow)
			if f != stats.Follow {
				continue
			}
		}
		res = append(res, stats)
	}
	resp.Resp(c, "100200", "成功", res)
}

func UpdateStats(c *gin.Context) {
	ch := "docker.container.stats"
	service.SendToAllServer(ch, map[string]interface{}{})
	resp.Resp(c, "100200", "成功", "")
}
