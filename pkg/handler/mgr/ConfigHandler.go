package mgr

import (
	"docker-manager/pkg/data"
	"docker-manager/pkg/data/table"
	"docker-manager/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func GetConfig(c *gin.Context) interface{} {
	conflist := data.GetConfigList()
	res := []table.Config{}
	for _, c := range conflist {
		if c.Name != "server.tokenSecret" {
			res = append(res, c)
		}
	}
	//boot.Resp(c, boot.CodeSuccess, "成功", res)
	return res
}

func UpdateConfig(c *gin.Context, req struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
	Memo  string `json:"Memo"`
}) {
	if req.Name != "agent.TaskFrequency" {
		boot.Resp(c, "100100", "配置参数错误", "")
		return
	}

	data.UpdateConfig(req.Name, req.Value, req.Memo, false)
	if req.Name == "agent.TaskFrequency" {
		ch := "base.config.update"
		service.SendToAllServer(ch, map[string]interface{}{})
	}

	boot.Resp(c, boot.CodeSuccess, "成功", "")
}

func GetWhiteList(c *gin.Context) {
	ipList, err := data.GetWhiteIpList()
	if err != nil {
		log.Println("GetWhiteList.err:", err)
	}
	boot.Resp(c, boot.CodeSuccess, "成功", ipList)
}

func AddWhiteIp(c *gin.Context, req struct {
	IP string `json:"IP"`
}) {
	data.AddWhiteIp(req.IP)
	service.LoadWhiteList()
	boot.Resp(c, boot.CodeSuccess, "成功", "")
}

func DeleteWhiteIp(c *gin.Context, req struct {
	IP string `json:"IP"`
}) {
	data.DelWhiteIp(req.IP)
	service.LoadWhiteList()
	boot.Resp(c, boot.CodeSuccess, "成功", "")
}
