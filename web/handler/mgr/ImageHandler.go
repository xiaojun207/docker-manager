package mgr

import (
	"docker-manager/data"
	"docker-manager/web/resp"
	"docker-manager/web/ws"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"log"
)

func ImageList(c *gin.Context) {
	res, err := data.GetImages()
	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	resp.Resp(c, "100200", "成功", res)
}

func ImageCmd(c *gin.Context) {
	//args := c.Query("args")
	//c.QueryArray("")
	args := []string{"docker", "start", "drone"}
	param := map[string]interface{}{
		"id":   uuid.New(),
		"args": args,
	}
	serverName := "docker-desktop"
	err := ws.AgentWsConnectGroup.Push(serverName, "base.cmd", param)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", "成功", err)
		return
	}
	resp.Resp(c, "100200", "成功", "")
}
