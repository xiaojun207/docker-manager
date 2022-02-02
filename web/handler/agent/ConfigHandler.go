package agent

import (
	"docker-manager/data"
	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) interface{} {
	return data.GetAgentConfig()
}
