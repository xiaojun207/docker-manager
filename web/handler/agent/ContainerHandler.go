package agent

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
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

	data.AddServer(table.Server{
		Name:            Name,
		OSType:          json["OSType"].(string),
		OperatingSystem: json["OperatingSystem"].(string),
		KernelVersion:   json["KernelVersion"].(string),
		DockerVersion:   json["ServerVersion"].(string),
		Running:         int(json["ContainersRunning"].(float64)),
		Containers:      int(json["Containers"].(float64)),
		Cpu:             int(json["NCPU"].(float64)),
		Memory:          int64(json["MemTotal"].(float64)),
		//PrivateIp:  "",
		//PublicIp:   "",
		//State:  "",
		Summary: utils2.MapToJson(json),
	})
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

	id := json["ID"].(string)
	Name := json["Name"].(string)

	for _, t := range json["conainers"].([]interface{}) {
		v := t.(map[string]interface{})
		v["AppId"] = AppId
		v["ServerName"] = Name
		v["Update"] = time.Now().Unix()
		ContainerName := utils.TrimContainerName(v["Names"])
		v["Name"] = ContainerName

		data.AddReplicas(ContainerName, Name)

		volumes := []string{}
		for _, m := range v["Mounts"].([]interface{}) {
			m := m.(map[string]interface{})
			vol := m["Source"].(string) + ":" + m["Destination"].(string)
			volumes = append(volumes, vol)
		}

		service := table.Service{
			Name:     ContainerName,
			Image:    v["Image"].(string),
			Ports:    utils2.MapToJson(v["Ports"]),
			Vol:      utils2.MapToJson(volumes),
			Running:  0,
			Replicas: 0,
		}
		data.AddService(service)

		var container table.Container
		utils2.MapToStruct(v, &container)
		container.ContainerId = v["Id"].(string)
		container.Summary = utils2.MapToJson(v)

		data.AddContainer(container)
	}
	//data.Container.Store(Name, json["conainers"])

	resp.Resp(c, "100200", "成功", gin.H{
		"id": id,
	})
}
