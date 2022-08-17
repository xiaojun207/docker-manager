package mgr

import (
	"docker-manager/pkg/data"
	"docker-manager/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func ExecClose(c *gin.Context, req struct {
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
	boot.Resp(c, "100200", "成功", "")
}
