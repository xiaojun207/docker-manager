package mgr

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/model"
	"docker-manager/service"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func ServiceList(c *gin.Context, page *model.Page) {
	res, err := data.GetServices(page)
	if err != nil {
		boot.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	boot.Resp(c, "100200", "成功", gin.H{
		"list": res,
		"page": page,
	})
}

func AppGroupList(c *gin.Context, page *model.Page) {
	res, err := data.GetServiceReplicas(page)
	if err != nil {
		boot.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	boot.Resp(c, "100200", "成功", gin.H{
		"list": res,
		"page": page,
	})
}

func DeleteGroup(c *gin.Context, req struct {
	Id int `json:"Id"`
}) {
	data.DeleteReplicas(req.Id)
	boot.Resp(c, "100200", "成功", "")
}

func DeleteService(c *gin.Context, req struct {
	Name string `json:"Name"`
}) {
	data.DeleteService(req.Name)
	boot.Resp(c, "100200", "成功", "")
}

func UpdateService(c *gin.Context, req table.Service) {
	data.AddService(req)
	log.Println("UpdateService.serviceInfo:", utils.StructToJson(req))
	serverNames := []string{""}
	service.PublishAppTask(serverNames, req)
	boot.Resp(c, "100200", "成功", "")
}
