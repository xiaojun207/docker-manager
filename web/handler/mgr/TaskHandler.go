package mgr

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/dto"
	"docker-manager/model"
	"docker-manager/service"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/gin-boot/boot"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func ContainerOperatorHandler(c *gin.Context, req struct {
	ServerNames []string `json:"ServerNames"`
	ContainerId string   `json:"ContainerId"`
}) {
	operator := c.Param("operator") // stop, remove, restart
	for _, serverName := range req.ServerNames {
		param := map[string]interface{}{
			"taskId":      uuid.New(),
			"containerId": req.ContainerId,
		}
		ch := "docker.container." + operator
		if operator == "remove" {
			service.DeleteContainer(req.ContainerId)
		}

		err := service.SaveAndSendTask(serverName, ch, param)
		if err != nil {
			log.Println(err)
			boot.Resp(c, "100100", "命令下发错误: "+err.Error(), "")
			return
		}
	}
	boot.Resp(c, "100200", "成功", "")
}

func ImageCmd(c *gin.Context, req struct {
	ServerName string `json:"serverName"`
	ImageId    string `json:"ImageId"`
}) {
	operator := c.Param("operator") // stop, remove, restart
	param := map[string]interface{}{
		"taskId":  uuid.New(),
		"imageId": req.ImageId,
	}
	ch := "docker.image." + operator

	err := service.SaveAndSendTask(req.ServerName, ch, param)
	if err != nil {
		log.Println(err)
		boot.Resp(c, "100100", "命令下发错误: "+err.Error(), "")
		return
	}
	boot.Resp(c, "100200", "成功", "")
}

func RePublishHandler(c *gin.Context, req struct {
	ServiceName string `json:"ServiceName"`
}) {
	appInfo, err := data.GetService(req.ServiceName)
	log.Println(appInfo, err)
}

func PublishYamlHandler(c *gin.Context, req struct {
	ServerNames []string `json:"serverNames"`
	Yaml        string   `json:"yaml"`
}) {
	service.PublishYaml(req.ServerNames, req.Yaml)
	boot.Resp(c, "100200", "成功", "")
}

func PublishHandler(c *gin.Context, serviceInfo dto.ServiceInfo) {
	s := table.Service{
		Name:     serviceInfo.Name,
		Image:    serviceInfo.Image,
		Envs:     serviceInfo.Env,
		Volumes:  serviceInfo.Volumes,
		Memory:   serviceInfo.Memory,
		Running:  serviceInfo.Running,
		Replicas: len(serviceInfo.ServerNames),
		Cover:    serviceInfo.Cover,
	}
	utils2.StructToMap(serviceInfo.Ports, &s.Ports)

	err := service.PublishAppTask(serviceInfo.ServerNames, s)
	if err != nil {
		log.Println(err)
		boot.Resp(c, "100100", err.Error(), "")
		return
	}
	boot.Resp(c, "100200", "成功", "")
}

func GetTasks(c *gin.Context, page model.Page) {
	list, _ := data.GetTasks(&page)
	boot.Resp(c, "100200", "成功", gin.H{
		"list": list,
		"page": page,
	})
}

func DelTask(c *gin.Context, req struct {
	Id int `json:"id"`
}) {
	data.DelTask(req.Id)
}
