package base

import (
	"docker-manager/data/base"
	"docker-manager/service"
	"docker-manager/web/resp"
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

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
	resp.Resp(c, "100200", "成功", data)
}

func ForbiddenLogHandler(c *gin.Context) {
	resp.Resp(c, "100200", "成功", service.ForbiddenLogMap())
}
