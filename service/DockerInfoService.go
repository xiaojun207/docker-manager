package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
)

// docker server注册记录
func DockerReg(json map[string]interface{}) error {
	//log.Println("DockerReg:", json)
	server := table.Server{
		Name:            json["Name"].(string),
		OSType:          json["OSType"].(string),
		OperatingSystem: json["OperatingSystem"].(string),
		KernelVersion:   json["KernelVersion"].(string),
		DockerVersion:   json["ServerVersion"].(string),
		Running:         int(json["ContainersRunning"].(float64)),
		Containers:      int(json["Containers"].(float64)),
		Cpu:             int(json["NCPU"].(float64)),
		Memory:          int64(json["MemTotal"].(float64)),
		Images:          int(json["Images"].(float64)),
		PrivateIp:       json["PrivateIp"].(string),
		HostIp:          json["HostIp"].(string),
		PublicIp:        json["PublicIp"].(string),
		//State:  "",
		Summary: json,
	}
	err := data.AddServer(server)
	return err
}
