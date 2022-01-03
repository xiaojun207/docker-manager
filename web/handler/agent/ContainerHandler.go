package agent

import (
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
)

func ContainersHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	AppId := c.GetHeader("AppId")

	id := json["ID"].(string)

	service.UpdateServerContainer(AppId, json)

	resp.Resp(c, "100200", "成功", gin.H{
		"id": id,
	})
}
