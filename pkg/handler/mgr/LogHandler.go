package mgr

import (
	"docker-manager/pkg/data"
	"docker-manager/pkg/model"
	"docker-manager/pkg/service"
	"docker-manager/pkg/ws"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func GetFollowLogList(c *gin.Context) {
	statsList, _ := data.GetContainerStatsList("true", []string{}, []string{}, &model.Page{})

	var res []map[string]interface{}
	for _, stats := range statsList {
		res = append(res, map[string]interface{}{
			"ID":         stats.ContainerId,
			"Name":       stats.Name,
			"ServerName": stats.ServerName,
		})
	}
	boot.Resp(c, boot.CodeSuccess, "成功", res)
}

func LogFollowStart(c *gin.Context, req struct {
	ContainerId string `json:"containerId"`
}) {
	param := map[string]interface{}{
		"taskId":      uuid.New(),
		"containerId": req.ContainerId,
	}
	serverName := data.GetServerNameByContainerId(req.ContainerId)
	if serverName == "" {
		log.Println(req.ContainerId + " server is not exists")
	}
	ch := "docker.container.log.follow"
	if !ws.AgentConnected(serverName) {
		boot.Resp(c, "100100", "服务器已离线", "")
		return
	}

	err := service.SaveAndSendTask(serverName, ch, param)
	log.Println("err:", err)
	boot.Resp(c, boot.CodeSuccess, "成功", "")
}

// 关闭日志流
func LogFollowClose(c *gin.Context, req struct {
	ContainerId string `json:"containerId"`
}) {
	param := map[string]interface{}{
		"taskId":      uuid.New(),
		"containerId": req.ContainerId,
	}
	serverName := data.GetServerNameByContainerId(req.ContainerId)

	ch := "docker.container.log.follow.close"
	err := service.SaveAndSendTask(serverName, ch, param)
	log.Println("err:", err)
	boot.Resp(c, boot.CodeSuccess, "成功", "")
}
