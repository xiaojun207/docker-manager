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
	ContainerName := json["ContainerName"].(string)
	appInfo := data.AppInfos.LoadMap(ContainerName)
	log.Println(appInfo)
}

func PublishHandler(c *gin.Context) {
	appInfo := dto.AppInfo{} //注意该结构接受的内容
	c.BindJSON(&appInfo)
	log.Println("appInfo:", appInfo)

	err := service.PublishAppTask(appInfo)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", err.Error(), "")
		return
	}
	resp.Resp(c, "100200", "成功", "")
}

func GetTasks(c *gin.Context) {
	tmps := data.Task.ValuesMap()
	utils.SortArrMap(tmps, func(a, b map[string]interface{}) bool {
		return a["ts"].(float64) > b["ts"].(float64)
	})
	resp.Resp(c, "100200", "成功", tmps)
}
