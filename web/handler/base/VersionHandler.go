package base

import (
	"docker-manager/service"
	"docker-manager/web/resp"
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//go:embed version
var Version string

func init() {
	log.Println("Version:", Version)
}

func VersionHandler(c *gin.Context) {
	latest := service.GetLatestTag()
	log.Println("Version Handler, current:", Version, ",latest:", latest)
	resp.Resp(c, "100200", "成功", gin.H{
		"current": Version,
		"latest":  service.GetLatestTag(),
	})
}

func VersionTextHandler(c *gin.Context) {
	c.JSON(http.StatusOK, []string{"1.2.3", "1.1.1"})
}
