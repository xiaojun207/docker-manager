package dto

type ContainerCreateConfig struct {
	ServerNames   []string        `json:"serverNames"`
	TaskId        string          `json:"taskId"`
	ImageName     string          `json:"imageName"`
	ContainerName string          `json:"containerName"`
	Ports         []ContainerPort `json:"ports"`
	Env           []string        `json:"env"`
	Volumes       []string        `json:"volumes"`
	Mounts        []struct {
		Destination string `json:"Destination"`
		Mode        string `json:"Mode"`
		Propagation string `json:"Propagation"`
		RW          bool   `json:"RW"`
		Source      string `json:"Source"`
		Type        string `json:"Type"`
	} `json:"Mounts"`
	Memory       int64             `json:"memory"`
	LogType      string            `json:"logType"`
	LogConfigMap map[string]string `json:"logConfig"`
}
