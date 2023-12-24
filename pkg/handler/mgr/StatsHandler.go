package mgr

import (
	"docker-manager/pkg/data"
	"docker-manager/pkg/model"
	"docker-manager/pkg/service"
	"docker-manager/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func GetStatsList(c *gin.Context, req struct {
	ServerNames    []string `json:"serverNames" form:"serverNames[]"`
	ContainerNames []string `json:"ContainerNames" form:"ContainerNames[]"`
	Follow         string   `json:"Follow" form:"Follow"`
}, page model.Page) {
	log.Println("GetStatsList:", req)

	var res []map[string]interface{}
	statss, _ := data.GetContainerStatsList(req.Follow, req.ServerNames, req.ContainerNames, &page)
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
	boot.Resp(c, boot.CodeSuccess, "成功", gin.H{
		"list": res,
		"page": page,
	})
}

func GetStats(c *gin.Context, req struct {
	ContainerId string `json:"ContainerId" form:"ContainerId"`
}) {
	res, _ := data.GetContainerStatss(req.ContainerId)
	boot.Resp(c, boot.CodeSuccess, "成功", res)
}

func UpdateStats(c *gin.Context) {
	ch := "docker.containers.stats"
	service.SendToAllServer(ch, map[string]interface{}{})
	boot.Resp(c, boot.CodeSuccess, "成功", "")
}
