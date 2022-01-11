package agent

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
)

func ImagesHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	AppId := c.GetHeader("AppId")

	//log.Printf("%v",&json)
	id := json["ID"].(string)
	Name := json["Name"].(string)

	for _, v := range json["Image"].([]interface{}) {
		v.(map[string]interface{})["AppId"] = AppId
		v.(map[string]interface{})["ServerName"] = Name
		var i table.Image
		utils2.MapToStruct(v.(map[string]interface{}), &i)
		data.AddImage(i)
	}

	boot.Resp(c, "100200", "成功", gin.H{
		"id": id,
	})
}
