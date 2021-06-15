package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = "100200"
)

type ApiResp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Resp(c *gin.Context, code string, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
