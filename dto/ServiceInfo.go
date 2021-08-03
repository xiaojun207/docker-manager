package dto

import "strings"

type ContainerPort struct {
	PrivatePort string `json:"PrivatePort"`
	Type        string `json:"Type"`
	IP          string `json:"IP"`
	PublicPort  string `json:"PublicPort"`
}

type ServiceInfo struct {
	ServerNames []string          `json:"ServerNames"`
	Image       string            `json:"Image"`
	Name        string            `json:"Name"`
	Ports       []ContainerPort   `json:"Ports"`
	Volumes     []string          `json:"Volumes"`
	Env         []string          `json:"Env"`
	Memory      int64             `json:"Memory"`
	LogConfig   map[string]string `json:"LogConfig"`
	LogType     string            `json:"LogType"`
	Running     int               `json:"Running"`
	Replicas    int               `json:"Replicas"`
}

func (e ServiceInfo) VolumeToArrMap() []map[string]interface{} {
	// [{"Destination":"/usr/local/openresty/nginx/conf/ssl","Mode":"","Propagation":"rprivate","RW":true,"Source":"/host_mnt/Users/kdaxrobot/IdeaProjects/local-openresty/nginx/conf/ssl","Type":"bind"},{"Destination":"/usr/local/openresty/nginx/logs","Mode":"","Propagation":"rprivate","RW":true,"Source":"/host_mnt/Users/kdaxrobot/IdeaProjects/local-openresty/logs","Type":"bind"},{"Destination":"/usr/local/openresty/nginx/tmp","Mode":"","Propagation":"rprivate","RW":true,"Source":"/host_mnt/Users/kdaxrobot/IdeaProjects/local-openresty/nginx/tmp","Type":"bind"},{"Destination":"/usr/local/openresty/extr_lua","Mode":"","Propagation":"rprivate","RW":true,"Source":"/host_mnt/Users/kdaxrobot/IdeaProjects/local-openresty/extr_lua","Type":"bind"},{"Destination":"/usr/local/openresty/nginx/conf/conf.d","Mode":"","Propagation":"rprivate","RW":true,"Source":"/host_mnt/Users/kdaxrobot/IdeaProjects/local-openresty/nginx/conf/conf.d","Type":"bind"},{"Destination":"/hostip","Mode":"ro","Propagation":"rprivate","RW":false,"Source":"/proc","Type":"bind"},{"Destination":"/usr/local/openresty/nginx/conf/lua","Mode":"","Propagation":"rprivate","RW":true,"Source":"/host_mnt/Users/kdaxrobot/IdeaProjects/local-openresty/nginx/conf/lua","Type":"bind"},{"Destination":"/usr/local/openresty/nginx/conf/nginx.conf","Mode":"","Propagation":"rprivate","RW":true,"Source":"/host_mnt/Users/kdaxrobot/IdeaProjects/local-openresty/nginx/conf/nginx.conf","Type":"bind"},{"Destination":"/data/web","Mode":"","Propagation":"rprivate","RW":true,"Source":"/host_mnt/Users/kdaxrobot/IdeaProjects/local-openresty/web","Type":"bind"},{"Destination":"/etc/hosts","Mode":"","Propagation":"rprivate","RW":true,"Source":"/host_mnt/Users/kdaxrobot/IdeaProjects/local-openresty/hosts","Type":"bind"}]
	var vol []map[string]interface{}
	for _, s2 := range e.Volumes {
		arr := strings.Split(s2, ":")
		e := map[string]interface{}{}
		e["Source"] = arr[0]
		e["Destination"] = arr[1]
		if len(arr) == 2 {
			e["RW"] = true
		} else {
			e["RW"] = arr[2] == "rw" //
		}
		vol = append(vol, e)
	}
	return vol
}

func (e ServiceInfo) EnvToArrMap() []map[string]interface{} {
	// xiaojun=123456
	var envs []map[string]interface{}
	for _, s2 := range e.Env {
		arr := strings.Split(s2, "=")
		e := map[string]interface{}{}
		e[arr[0]] = arr[1]
		envs = append(envs, e)
	}
	return envs
}
