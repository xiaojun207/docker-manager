package mgr

import (
	"docker-manager/data"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
)

func ImageList(c *gin.Context) {
	serverName := c.Query("serverName")
	tagName := c.Query("tagName")
	res, err := data.GetImages(serverName, tagName)
	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	resp.Resp(c, "100200", "成功", res)
}
