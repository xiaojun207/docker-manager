package main

import (
	"docker-manager/conf"
	"docker-manager/data"
	"docker-manager/service"
	"docker-manager/web"
	"github.com/xiaojun207/go-base-utils/utils"
)

func main() {
	conf.ParseParam()

	data.InitDB(conf.DriverName, conf.DataSourceUrl, utils.StrToBool(conf.UseCache))
	service.InitTokenHelper()
	service.LoadWhiteList()
	service.LoadContainerMap()

	web.Start(conf.Port, conf.ContextPath)

	defer data.Close()

	select {}
}
