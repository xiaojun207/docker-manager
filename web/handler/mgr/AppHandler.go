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

func AppList(c *gin.Context) {
	res := data.AppInfos.ValuesMap()
	resp.Resp(c, "100200", "成功", res)
}

func AppGroupList(c *gin.Context) {
	res := data.AppGroup.ToStrMap()
	resp.Resp(c, "100200", "成功", res)
}

func DelApp(c *gin.Context) {
	appInfo := dto.AppInfo{} //注意该结构接受的内容
	c.BindJSON(&appInfo)
	data.AppInfos.Remove(appInfo.Name)
	resp.Resp(c, "100200", "成功", "")
}

func UpdateApp(c *gin.Context) {
	appInfo := dto.AppInfo{} //注意该结构接受的内容
	c.BindJSON(&appInfo)
	data.GetAppInfo(appInfo.Name, &appInfo)
	log.Println("UpdateApp.appInfo:", utils.StructToJson(appInfo))
	service.PublishAppTask(appInfo)
	resp.Resp(c, "100200", "成功", "")
}
