package mgr

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/model"
	"docker-manager/service"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func ServiceList(c *gin.Context) {
	page := model.GetPage(c)
	res, err := data.GetServices(&page)
	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	resp.Resp(c, "100200", "成功", gin.H{
		"list": res,
		"page": page,
	})
}

func AppGroupList(c *gin.Context) {
	page := model.GetPage(c)
	res, err := data.GetServiceReplicas(&page)
	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	resp.Resp(c, "100200", "成功", gin.H{
		"list": res,
		"page": page,
	})
}

func DeleteGroup(c *gin.Context) {
	replicas := table.ServiceReplicas{} //注意该结构接受的内容
	c.BindJSON(&replicas)
	data.DeleteReplicas(replicas.Id)
	resp.Resp(c, "100200", "成功", "")
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
