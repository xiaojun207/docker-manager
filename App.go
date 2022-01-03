package main

import (
	"docker-manager/data"
	"docker-manager/service"
	"docker-manager/web"
	"flag"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

var (
	driverName    = "sqlite3"
	dataSourceUrl = "data/database.db"
	useCache      = "true"
)

func ParseParam() {
	flag.StringVar(&driverName, "driverName", driverName, "driverName, eg.: mysql or sqlite3")
	flag.StringVar(&useCache, "useCache", useCache, "useCache, eg.: true or false")
	flag.StringVar(&dataSourceUrl, "dataSourceUrl", dataSourceUrl, "dataSourceUrl, eg.: root:123456@(127.0.0.1:3306)/docker-manager?charset=utf8")
	flag.Parse()
	log.Println("flag.driverName:", driverName)
	log.Println("flag.useCache:", useCache)
	log.Println("flag.dataSourceUrl:", utils.SubstrAfter(dataSourceUrl, "@"))
}

func main() {
	ParseParam()

	data.InitDB(driverName, dataSourceUrl, utils.StrToBool(useCache))
	service.InitTokenHelper()
	service.LoadWhiteList()
	service.LoadContainerMap()
	service.LoadContainerStatsMap()
	web.Start("8068", "/dockerMgrApi/")

	defer data.Close()

	select {}
}
