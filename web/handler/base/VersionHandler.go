package base

import (
	"docker-manager/service"
	_ "embed"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
	"net/http"
	"strings"
)

//go:embed version
var Version string

func init() {
	Version = strings.TrimSpace(Version)
	log.Println("Version:", Version)
}

func VersionHandler(c *gin.Context) {
	latest := service.GetLatestTag()
	log.Println("Version Handler, current:", Version, ",latest:", latest)
	boot.Resp(c, "100200", "成功", gin.H{
		"current": Version,
		"latest":  latest,
	})
}

func VersionTextHandler(c *gin.Context) {
	c.JSON(http.StatusOK, []string{"1.2.3", "1.1.1"})
}

func TestHandler(c *gin.Context) interface{} {
	key := c.Query("key")
	if key == "1" {
		return boot.NewWebError("100103", "错误1")
	} else if key == "2" {
		return errors.New("错误2")
	} else if key == "3" {
		return boot.ApiResp{"100103", "字段校验错误", "错误3"}
	}

	return "this is a test!"
}
