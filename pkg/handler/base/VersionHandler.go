package base

import (
	"docker-manager/pkg/service"
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
)

func VersionHandler(c *gin.Context) interface{} {
	latest := service.GetLatestTag()
	log.Println("Version Handler, current:", service.Version, ",latest:", latest)
	return gin.H{
		"current": service.Version,
		"latest":  latest,
	}
}
