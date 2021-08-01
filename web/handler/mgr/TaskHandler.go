package mgr

import (
	"docker-manager/data"
	"docker-manager/dto"
	"docker-manager/service"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"log"
)

func ContainerOperatorHandler(c *gin.Context) {
	operator := c.Param("operator")      // stop, remove, restart
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	log.Println("json:", json)
	serverNames := utils.ArrInterfaceToStr(json["serverNames"].([]interface{}))
	containerId := json["containerId"].(string)

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

	err := service.PublishAppTask(serviceInfo)
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
