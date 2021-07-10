package main

import (
	"docker-manager/data"
	"docker-manager/web"
	"flag"
	"log"
)

var (
	driverName    string
	dataSourceUrl string
)

func ParseParam() {
	flag.StringVar(&driverName, "driverName", driverName, "driverName, eg.: mysql or sqlite3")
	flag.StringVar(&dataSourceUrl, "dataSourceUrl", dataSourceUrl, "dataSourceUrl, eg.: root:123456@(127.0.0.1:3306)/docker-manager?charset=utf8")
	flag.Parse()
	log.Println("flag.driverName:", driverName)
	log.Println("flag.dataSourceUrl:", dataSourceUrl)
}

func main() {
	//ParseParam()

	//data.InitDB(driverName, dataSourceUrl)
	web.Start("8068", "/dockerMgrApi/")

	defer data.Close()

	select {}
}
