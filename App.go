package main

import (
	"docker-manager/pkg/conf"
	"docker-manager/pkg/data"
	"docker-manager/pkg/routers"
	"docker-manager/pkg/service"
	"github.com/xiaojun207/gin-boot/boot"
	"github.com/xiaojun207/go-base-utils/utils"
)

func main() {
	conf.ParseParam()

	data.InitDB(conf.DriverName, conf.DataSourceUrl, utils.StrToBool(conf.UseCache))
	service.InitTokenHelper()
	service.LoadWhiteList()
	service.LoadContainerMap()

	boot.Start(conf.Port, "/", func(webRouter *boot.WebRouter) {
		routers.StaticRouter(webRouter)
		routers.ApiRouter(webRouter.Group(conf.ContextPath))
	})

	defer data.Close()

	select {}
}
