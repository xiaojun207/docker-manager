package main

import (
	"docker-manager/data"
	"docker-manager/web"
	"flag"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

var (
	driverName    = "mysql"
	dataSourceUrl string
	useCache      = false
)

func ParseParam() {
	flag.StringVar(&driverName, "driverName", driverName, "driverName, eg.: mysql or sqlite3")
	flag.BoolVar(&useCache, "useCache", useCache, "useCache, eg.: true or false")
	flag.StringVar(&dataSourceUrl, "dataSourceUrl", dataSourceUrl, "dataSourceUrl, eg.: root:123456@(127.0.0.1:3306)/docker-manager?charset=utf8")
	flag.Parse()
	log.Println("flag.driverName:", driverName)
	log.Println("flag.useCache:", useCache)
	log.Println("flag.dataSourceUrl:", utils.SubstrAfter(dataSourceUrl, "@"))
}

func main() {
	ParseParam()

	data.InitDB(driverName, dataSourceUrl, useCache)

	web.Start("8068", "/dockerMgrApi/")

	defer data.Close()

	select {}
}
