package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func PublishAppTask(serverNames []string, service table.Service) error {
	log.Println("PublishAppTask.appInfo:", utils.StructToJson(service))
	data.AddService(service)

	for _, serverName := range serverNames {
		param := map[string]interface{}{
			"taskId":        uuid.New(),
			"serverName":    serverName,
			"imageName":     service.Image,
			"containerName": service.Name,
			"ports":         service.Ports,
			"volumes":       service.Volumes,
			"env":           service.Envs, // {"appversion=1.0.1"}
			"memory":        service.Memory,
			"cover":         service.Cover,
			//"logType":       service.LogType,
			//"logConfig":     service.LogConfig,
		}
		ch := "docker.container.run"

		err := SaveAndSendTask(serverName, ch, param)
		log.Println("SaveAndSendTask.err:", err)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func PublishYaml(serverNames []string, yamlStr string) {
	configFile := types.ConfigFile{
		Content: []byte(yamlStr),
	}
	project, err := loader.Load(types.ConfigDetails{
		ConfigFiles: append([]types.ConfigFile{}, configFile),
	})
	if err != nil {
		log.Println(err)
	}

	serer := project.Services[0]
	//j, _ := json.Marshal(serer)
	//log.Println(string(j))
	s := ToServiceTable(serer)
	data.AddService(s)

	for _, serverName := range serverNames {
		param := map[string]interface{}{
			"taskId":        uuid.New(),
			"serverName":    serverName,
			"imageName":     s.Image,
			"containerName": s.Name,
			"ports":         s.Ports,
			"volumes":       s.Volumes,
			"env":           s.Envs, // {"appversion=1.0.1"}
			//"memory":        s.Memory,
			"cover": s.Cover,
			//"logType":       service.LogType,
			//"logConfig":     service.LogConfig,
		}
		ch := "docker.container.run"

		err := SaveAndSendTask(serverName, ch, param)
		log.Println("SaveAndSendTask.err:", err)
		if err != nil {
			log.Println(err)
			return
		}
	}

}

func ToServiceTable(serer types.ServiceConfig) table.Service {
	var env []string
	for k, v := range serer.Environment {
		log.Println("Environment:", k, *v)
		env = append(env, k+":"+*v)
	}
	var volumes []string
	for _, v := range serer.Volumes {
		volumes = append(volumes, v.Source+":"+v.Target+":rw")
	}
	var ports []map[string]interface{}
	for _, port := range serer.Ports {
		ports = append(ports, map[string]interface{}{
			"IP":          "0.0.0.0",
			"Type":        port.Protocol,
			"PrivatePort": utils.Uint64ToStr(uint64(port.Target)),
			"PublicPort":  utils.Uint64ToStr(utils.StrToUint64(port.Published)),
		})
		// {"IP":"0.0.0.0","PrivatePort":"80","PublicPort":"80","Type":"tcp"}
	}
	return table.Service{
		Name:     serer.ContainerName,
		Image:    serer.Image,
		Ports:    ports,
		Envs:     env,
		Volumes:  volumes,
		Memory:   int64(serer.MemLimit),
		Running:  1,
		Replicas: 1,
		Cover:    true,
	}
}
