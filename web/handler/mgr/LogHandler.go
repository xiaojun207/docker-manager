package mgr

import (
	"docker-manager/data"
	"docker-manager/service"
	"docker-manager/web/resp"
	"docker-manager/web/ws"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"log"
)

func GetFollowLogList(c *gin.Context) {
	statsList, _ := data.GetContainerStatsList([]string{})

	var res []map[string]interface{}
	for _, stats := range statsList {
		if stats.Follow {
			res = append(res, map[string]interface{}{
				"ID":         stats.ContainerId,
				"Name":       stats.Name,
				"ServerName": stats.ServerName,
			})
		}
	}
	resp.Resp(c, "100200", "成功", res)
}

func LogFollowStart(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	containerId := json["containerId"].(string)
	param := map[string]interface{}{
		"taskId":      uuid.New(),
		"containerId": containerId,
	}
	serverName := data.GetServerNameByContainerId(containerId)
	if serverName == "" {
		log.Println(containerId + " server is not exists")
	}
	ch := "docker.container.log.follow"
	if !ws.AgentConnected(serverName) {
		resp.Resp(c, "100100", "服务器已离线", "")
		return
	}

	err := service.SaveAndSendTask(serverName, ch, param)
	log.Println("err:", err)
	resp.Resp(c, "100200", "成功", "")
}

func LogFollowClose(c *gin.Context) {
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
