package mgr

import (
	"docker-manager/data"
	"docker-manager/dto"
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func ServiceList(c *gin.Context) {
	res, _ := data.GetServices()
	resp.Resp(c, "100200", "成功", res)
}

func AppGroupList(c *gin.Context) {
	res, _ := data.GetServiceReplicas()
	resp.Resp(c, "100200", "成功", res)
}

func DeleteService(c *gin.Context) {
	serviceInfo := dto.ServiceInfo{} //注意该结构接受的内容
	c.BindJSON(&serviceInfo)
	data.DeleteService(serviceInfo.Name)
	resp.Resp(c, "100200", "成功", "")
}

func UpdateService(c *gin.Context) {
	serviceInfo := dto.ServiceInfo{} //注意该结构接受的内容
	c.BindJSON(&serviceInfo)
	data.AddAppInfo(serviceInfo.Name, serviceInfo)
	log.Println("UpdateService.serviceInfo:", utils.StructToJson(serviceInfo))
	service.PublishAppTask(serviceInfo)
	resp.Resp(c, "100200", "成功", "")
}
