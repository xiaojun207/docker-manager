package mgr

import (
	"docker-manager/data"
	"docker-manager/model"
	"docker-manager/service"
	"docker-manager/utils"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func GetStatsList(c *gin.Context) {
	serverNames := c.QueryArray("serverNames[]")
	ContainerNames := c.QueryArray("ContainerNames[]")
	Follow := c.Query("Follow")
	page := model.GetPage(c)
	log.Println("serverNames:", serverNames, ",ContainerNames:", ContainerNames, ",Follow:", Follow)

	var res []map[string]interface{}
	statss, _ := data.GetContainerStatsList(Follow, serverNames, ContainerNames, &page)
	for _, stats := range statss {
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
	boot.Resp(c, "100200", "成功", gin.H{
		"list": res,
		"page": page,
	})
}

func GetStats(c *gin.Context) {
	ContainerId := c.Query("ContainerId")
	res, _ := data.GetContainerStatss(ContainerId)
	boot.Resp(c, "100200", "成功", res)
}

func UpdateStats(c *gin.Context) {
	ch := "docker.containers.stats"
	service.SendToAllServer(ch, map[string]interface{}{})
	boot.Resp(c, "100200", "成功", "")
}
