package mgr

import (
	"docker-manager/data"
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) {
	resp.Resp(c, "100200", "成功", data.Config.ToStrMap())
}

func UpdateConfig(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	TaskFrequency := json["TaskFrequency"]

	data.Config.Store("agentConfig", map[string]interface{}{
		"TaskFrequency": TaskFrequency,
	})

	ch := "base.config.update"
	service.SendToAllServer(ch, map[string]interface{}{})
	resp.Resp(c, "100200", "成功", "")
}
