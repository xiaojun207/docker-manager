package mgr

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) {
	conflist := data.GetConfigList()
	res := []table.Config{}
	for _, c := range conflist {
		if c.Name != "server.tokenSecret" {
			res = append(res, c)
		}
	}
	resp.Resp(c, "100200", "成功", res)
}

func UpdateConfig(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	name := json["Name"].(string)
	value := json["Value"].(string)
	memo := json["Memo"].(string)

	if name != "agent.TaskFrequency" {
		resp.Resp(c, "100100", "配置参数错误", "")
		return
	}

	data.UpdateConfig(name, value, memo, false)
	if name == "agent.TaskFrequency" {
		ch := "base.config.update"
		service.SendToAllServer(ch, map[string]interface{}{})
	}

	resp.Resp(c, "100200", "成功", "")
}
