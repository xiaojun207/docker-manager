package agent

import (
	"docker-manager/service"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
)

func ContainersHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	AppId := c.GetHeader("AppId")

	id := json["ID"].(string)

	service.UpdateServerContainer(AppId, json)

	boot.Resp(c, "100200", "成功", gin.H{
		"id": id,
	})
}
