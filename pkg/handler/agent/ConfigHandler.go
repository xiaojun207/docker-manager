package agent

import (
	"docker-manager/pkg/data"
	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) interface{} {
	return data.GetAgentConfig()
}
