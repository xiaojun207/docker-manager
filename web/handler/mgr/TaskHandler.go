package mgr

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/dto"
	"docker-manager/service"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func ContainerOperatorHandler(c *gin.Context) {
	operator := c.Param("operator")      // stop, remove, restart
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	log.Println("json:", json)
	serverNames := utils.ArrInterfaceToStr(json["ServerNames"].([]interface{}))
	containerId := json["ContainerId"].(string)

	for _, serverName := range serverNames {
		param := map[string]interface{}{
			"taskId":      uuid.New(),
			"containerId": containerId,
		}
		ch := "docker.container." + operator

		err := service.SaveAndSendTask(serverName, ch, param)
		if err != nil {
			log.Println(err)
			resp.Resp(c, "100100", "命令下发错误: "+err.Error(), "")
			return
		}
	}
	resp.Resp(c, "100200", "成功", "")
}

func ImageCmd(c *gin.Context) {
	operator := c.Param("operator")      // stop, remove, restart
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	log.Println("json:", json)
	serverName := json["serverName"].(string)
	ImageId := json["ImageId"].(string)

	param := map[string]interface{}{
		"taskId":  uuid.New(),
		"imageId": ImageId,
	}
	ch := "docker.image." + operator

	err := service.SaveAndSendTask(serverName, ch, param)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", "命令下发错误: "+err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", "")
}

func RePublishHandler(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	ServiceName := json["ServiceName"].(string)
	appInfo, err := data.GetService(ServiceName)
	log.Println(appInfo, err)
}

func PublishHandler(c *gin.Context) {
	serviceInfo := dto.ServiceInfo{} //注意该结构接受的内容
	c.BindJSON(&serviceInfo)
	log.Println("serviceInfo:", serviceInfo)
	s := table.Service{
		Name:     serviceInfo.Name,
		Image:    serviceInfo.Image,
		Env:      serviceInfo.EnvToArrMap(),
		Vol:      serviceInfo.VolumeToArrMap(),
		Memory:   serviceInfo.Memory,
		Running:  serviceInfo.Running,
		Replicas: len(serviceInfo.ServerNames),
	}
	utils2.StructToMap(serviceInfo.Ports, &s.Ports)

	err := service.PublishAppTask(serviceInfo.ServerNames, s)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", "")
}

func GetTasks(c *gin.Context) {
	tmps, _ := data.GetTasks()
	resp.Resp(c, "100200", "成功", tmps)
}
