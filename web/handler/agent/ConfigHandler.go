package agent

import (
	"docker-manager/data"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) {
	resp.Resp(c, "100200", "成功", data.GetAgentConfig())
}
