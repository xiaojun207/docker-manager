package agent

import (
	"docker-manager/data"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
)

func ImagesHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	AppId := c.GetHeader("AppId")

	//log.Printf("%v",&json)
	id := json["ID"].(string)
	Name := json["Name"].(string)

	for _, v := range json["Images"].([]interface{}) {
		v.(map[string]interface{})["AppId"] = AppId
		v.(map[string]interface{})["ServerName"] = Name
	}

	data.Images.Store(Name, json["Images"])

	resp.Resp(c, "100200", "成功", gin.H{
		"id": id,
	})
}
