package base

import (
	"docker-manager/data/base"
	"docker-manager/service"
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// GET /dbtrace
func DBTraceHandler(c *gin.Context) interface{} {
	log.Println("DBTraceHandler")
	timeLen := time.Now().Unix() - base.DBTracingHook.StartTime.Unix()

	data := map[string]interface{}{
		"AfterNum":  base.DBTracingHook.AfterNum,
		"BeforeNum": base.DBTracingHook.BeforeNum,
		"timeLen":   timeLen,
		"per":       base.DBTracingHook.BeforeNum / timeLen,
		"SqlMap":    base.DBTracingHook.SqlMap.ToStrMap(),
	}
	return data
}

func ForbiddenLogHandler(c *gin.Context) interface{} {
	return service.ForbiddenLogMap()
}
