package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/utils"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"time"
)

func UpdateServerContainer(AppId string, json map[string]interface{}) {
	Name := json["Name"].(string)

	containerMap := map[string]table.Container{}
	for _, t := range json["conainers"].([]interface{}) {
		v := t.(map[string]interface{})
		v["AppId"] = AppId
		v["ServerName"] = Name
		v["Update"] = time.Now().Unix()
		ContainerName := utils.TrimContainerName(v["Names"])
		v["Name"] = ContainerName

		data.AddReplicas(ContainerName, Name)

		var service = table.Service{
			Name:     ContainerName,
			Image:    v["Image"].(string),
			Vol:      utils.ArrInterfaceToMap(v["Mounts"].([]interface{})),
			Running:  0,
			Ports:    utils.ArrInterfaceToMap(v["Ports"].([]interface{})),
			Replicas: 0,
		}
		data.AddService(service)

		var container table.Container
		utils2.MapToStruct(v, &container)
		container.ContainerId = v["Id"].(string)
		container.Summary = v

		data.AddContainer(container)
		containerMap[ContainerName] = container
	}

	dbArr, _ := data.GetContainersByServerName(Name)
	for _, container := range dbArr {
		_, ok := containerMap[container.Name]
		if !ok {
			// 如果没有，说明已经删除了
			data.DelContainer(container)
		}
	}

}