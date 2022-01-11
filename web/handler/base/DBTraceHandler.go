package base

import (
	"docker-manager/data/base"
	"docker-manager/service"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
	"time"
)

// GET /dbtrace
func DBTraceHandler(c *gin.Context) {
	log.Println("DBTraceHandler")
	timeLen := time.Now().Unix() - base.DBTracingHook.StartTime.Unix()

	data := map[string]interface{}{
		"AfterNum":  base.DBTracingHook.AfterNum,
		"BeforeNum": base.DBTracingHook.BeforeNum,
		"timeLen":   timeLen,
		"per":       base.DBTracingHook.BeforeNum / timeLen,
		"SqlMap":    base.DBTracingHook.SqlMap.ToStrMap(),
	}
	boot.Resp(c, "100200", "成功", data)
}

func ForbiddenLogHandler(c *gin.Context) {
	boot.Resp(c, "100200", "成功", service.ForbiddenLogMap())
}
