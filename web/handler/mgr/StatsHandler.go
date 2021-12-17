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

func GetStatsList(c *gin.Context) {
	serverNames := c.QueryArray("serverNames[]")
	ContainerNames := c.QueryArray("ContainerNames[]")
	Follow := c.Query("Follow")
	log.Println("serverNames:", serverNames, ",ContainerNames:", ContainerNames, ",Follow:", Follow)

	var res []map[string]interface{}
	statss, _ := data.GetContainerStatsList(serverNames)
	for _, stats := range statss {
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

		res = append(res, map[string]interface{}{
			"Name":               stats.Name,
			"ServerName":         stats.ServerName,
			"CpuStats":           utils.FormatCpu(stats.CpuStats, stats.PrecpuStats),
			"MemoryStats":        utils.FormatMemory(stats.MemoryStats),
			"MemoryStatsPercent": utils.FormatMemoryPercent(stats.MemoryStats),
			"Networks":           utils.FormatNet(stats.Networks),
			"ContainerId":        stats.ContainerId,
		})
	}
	resp.Resp(c, "100200", "成功", res)
}

func GetStats(c *gin.Context) {
	ContainerId := c.Query("ContainerId")
	res, _ := data.GetContainerStatss(ContainerId)
	resp.Resp(c, "100200", "成功", res)
}

func UpdateStats(c *gin.Context) {
	ch := "docker.containers.stats"
	service.SendToAllServer(ch, map[string]interface{}{})
	resp.Resp(c, "100200", "成功", "")
}
