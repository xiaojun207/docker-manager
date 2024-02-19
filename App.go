package main

import (
	"docker-manager/pkg/conf"
	"docker-manager/pkg/data"
	"docker-manager/pkg/jobs"
	"docker-manager/pkg/routers"
	"docker-manager/pkg/service"

	"github.com/xiaojun207/gin-boot/boot"
	"github.com/xiaojun207/go-base-utils/utils"
)

func main() {
	conf.ParseParam()

	// 初始化数据库
	data.InitDB(conf.DriverName, conf.DataSourceUrl, utils.StrToBool(conf.UseCache))
	service.InitTokenHelper()
	service.LoadWhiteList()
	service.LoadContainerMap()

	boot.Start(conf.Port, "/", func(webRouter *boot.WebRouter) {
		routers.StaticRouter(webRouter)
		routers.ApiRouter(webRouter.Group(conf.ContextPath))
	})

	jobs.StartJobs()

	defer data.Close()

	select {}
}
