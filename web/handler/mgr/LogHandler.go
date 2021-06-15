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
	tmps := []map[string]interface{}{}
	data.Stats.ForEachArr(func(serverName string, arr []interface{}) {
		for _, value := range arr {
			m := value.(map[string]interface{})
			if m["Follow"].(bool) {
				c := map[string]interface{}{
					"ID":         m["ID"],
					"Name":       m["Name"],
					"ServerName": serverName,
				}
				tmps = append(tmps, c)
			}
		}
	})
	resp.Resp(c, "100200", "成功", tmps)
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
	if !ws.IsConnected(serverName) {
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
