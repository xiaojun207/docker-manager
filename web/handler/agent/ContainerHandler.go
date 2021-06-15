package agent

import (
	"docker-manager/data"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var (
	Version = "0.0.1"
)

func RegDockerHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	Name := json["Name"].(string)
	id := json["ID"].(string)
	//AppId := c.GetHeader("AppId")

	data.Servers.Store(Name, json)
	log.Println("reg server:", id, " Name:", Name)

	resp.Resp(c, "100200", "成功", gin.H{
		"name": Name,
		"id":   id,
	})
}

func ContainersHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	AppId := c.GetHeader("AppId")

	//log.Printf("%v",&json)
	id := json["ID"].(string)
	Name := json["Name"].(string)

	for _, t := range json["conainers"].([]interface{}) {
		v := t.(map[string]interface{})
		v["AppId"] = AppId
		v["ServerName"] = Name
		v["Update"] = time.Now().Unix()
		ContainerName := utils.TrimContainerName(v["Names"])
		v["Name"] = ContainerName

		data.ContainerServerMap.StoreStr(v["Id"].(string), Name)
		data.AddAppGroup(ContainerName, Name)

		volumes := []string{}
		for _, m := range v["Mounts"].([]interface{}) {
			m := m.(map[string]interface{})
			vol := m["Source"].(string) + ":" + m["Destination"].(string)
			volumes = append(volumes, vol)
		}
		appInfo := map[string]interface{}{
			"Name":    ContainerName,
			"Image":   v["Image"],
			"Ports":   v["Ports"],
			"Volumes": volumes,
		}
		data.AddAppInfo(ContainerName, appInfo)
	}

	data.Containers.Store(Name, json["conainers"])

	resp.Resp(c, "100200", "成功", gin.H{
		"id": id,
	})
}
