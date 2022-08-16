package base

import (
	"docker-manager/service"
	_ "embed"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
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

func VersionTextHandler(c *gin.Context) interface{} {
	return []string{"1.2.3", "1.1.1"}
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
