package dto

type ContainerPort struct {
	PrivatePort string `json:"PrivatePort"`
	Type        string `json:"Type"`
	IP          string `json:"IP"`
	PublicPort  string `json:"PublicPort"`
}

type AppInfo struct {
	ServerNames []string          `json:"ServerNames"`
	Image       string            `json:"Image"`
	Name        string            `json:"Name"`
	Ports       []ContainerPort   `json:"Ports"`
	Volumes     []string          `json:"Volumes"`
	Env         []string          `json:"Env"`
	Memory      int64             `json:"Memory"`
	LogConfig   map[string]string `json:"LogConfig"`
	LogType     string            `json:"LogType"`
}
