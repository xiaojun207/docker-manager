package mgr

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func ServiceList(c *gin.Context) {
	res, err := data.GetServices()
	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	resp.Resp(c, "100200", "成功", res)
}

func AppGroupList(c *gin.Context) {
	res, err := data.GetServiceReplicas()
	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	resp.Resp(c, "100200", "成功", res)
}

func DeleteService(c *gin.Context) {
	service := table.Service{} //注意该结构接受的内容
	c.BindJSON(&service)
	data.DeleteService(service.Name)
	resp.Resp(c, "100200", "成功", "")
}

func UpdateService(c *gin.Context) {
	serviceInfo := table.Service{} //注意该结构接受的内容
	c.BindJSON(&serviceInfo)
	data.AddService(serviceInfo)
	log.Println("UpdateService.serviceInfo:", utils.StructToJson(serviceInfo))
	serverNames := []string{""}
	service.PublishAppTask(serverNames, serviceInfo)
	resp.Resp(c, "100200", "成功", "")
}
