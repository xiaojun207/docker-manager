package mgr

import (
	"docker-manager/data"
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"log"
)

func ExecClose(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	containerId := json["containerId"].(string)
	param := map[string]interface{}{
		"taskId":      uuid.New(),
		"containerId": containerId,
	}
	serverName := data.GetServerNameByContainerId(containerId)

	ch := "docker.container.log.follow.close"
	err := service.SaveAndSendTask(serverName, ch, param)
	log.Println("err:", err)
	resp.Resp(c, "100200", "成功", "")
}
