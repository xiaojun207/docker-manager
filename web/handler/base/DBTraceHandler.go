package base

import (
	"docker-manager/data/base"
	"docker-manager/web/resp"
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
)

func DBTraceHandler(c *gin.Context) {
	log.Println("DBTraceHandler")

	data := map[string]interface{}{
		"AfterNum":  base.DBTracingHook.AfterNum,
		"BeforeNum": base.DBTracingHook.BeforeNum,
		"SqlMap":    base.DBTracingHook.SqlMap.ToStrMap(),
	}
	resp.Resp(c, "100200", "成功", data)
}
