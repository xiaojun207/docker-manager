package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/model"
	"docker-manager/utils"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
	"time"
)

var ContainerMap model.SyncMap
var ContainerShortIdServerNameMap model.SyncMap

func LoadContainerMap() {
	records, err := data.GetContainers()
	if err != nil {
		log.Println("LoadContainerIdMap.err:", err)
	}
	ContainerMap.RemoveAll()
	ContainerShortIdServerNameMap.RemoveAll()
	for _, record := range records {
		AddContainerMapCache(record)
	}
	log.Println("LoadContainerIdMap.size:", ContainerMap.Size())
}

func AddContainerMapCache(record table.Container) {
	ContainerMap.Store(record.ContainerId, record)
	ContainerShortIdServerNameMap.Store(utils.ContainerShortId(record.ContainerId), record.ServerName)
}

func removeContainerMapCache(containerId string) {
	ContainerMap.Remove(containerId)
	ContainerShortIdServerNameMap.Remove(utils.ContainerShortId(containerId))
}

func UpdateServerContainer(AppId string, json map[string]interface{}) {
	ServerName := json["Name"].(string)

	containerMap := map[string]table.Container{}
	for _, t := range json["containers"].([]interface{}) {
		v := t.(map[string]interface{})
		v["AppId"] = AppId
		v["ServerName"] = ServerName
		v["Update"] = time.Now().Unix()
		ContainerName := utils.TrimContainerName(v["Names"])
		v["Name"] = ContainerName

		data.AddReplicas(ContainerName, ServerName)

		var service = table.Service{
			Name:     ContainerName,
			Image:    v["Image"].(string),
			Volumes:  GetVolumes(v["Mounts"].([]interface{})),
			Running:  0,
			Ports:    utils.ArrInterfaceToMap(v["Ports"].([]interface{})),
			Replicas: 0,
		}
		data.AddService(service)

		var container table.Container
		utils2.MapToStruct(v, &container)
		container.ContainerId = v["Id"].(string)
		container.Summary = v

		data.AddContainer(&container)
		AddContainerMapCache(container)
		containerMap[ContainerName] = container
	}

	dbArr, _ := data.GetContainersByServerName(ServerName)
	for _, container := range dbArr {
		_, ok := containerMap[container.Name]
		if !ok {
			// 如果没有，说明已经删除了
			data.DelContainer(container)
			removeContainerMapCache(container.ContainerId)
		}
	}

}

func DeleteContainer(containerId string) {
	c, err := data.GetContainer(containerId)
	if err != nil {
		log.Println("DeleteContainer.containerId", containerId)
		return
	}
	data.DelContainer(c)
}

func GetVolumes(volmap []interface{}) (res []string) {
	for _, t := range volmap {
		v := t.(map[string]interface{})
		tmp := v["Source"].(string) + ":" + v["Destination"].(string) + ":" + utils2.If(v["RW"].(bool), "rw", "ro").(string)
		res = append(res, tmp)
	}
	return
}

func GetServerNameByContainerId(containerId string) string {
	return ContainerShortIdServerNameMap.GetStr(containerId)
}

func GetServerNameByContainerShortId(cId string) string {
	return ContainerShortIdServerNameMap.GetStr(cId)
}
