package agent

import (
	"docker-manager/data"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
)

func GetConfig(c *gin.Context) {
	boot.Resp(c, "100200", "成功", data.GetAgentConfig())
}
