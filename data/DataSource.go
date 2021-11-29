package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	"log"
	//"docker-manager/model"
)

func Close() {
	base.CloseDBEngine()
}

func InitDB(driverName, dataSourceUrl string, useCache bool) {
	//Init("sqlite3", "./db/data.db")
	//InitDB("mysql", "root:Abc123@(127.0.0.1:3306)/docker-manager?charset=utf8")
	//DBEngine.Sync2(new(table.User), new(table.UserApi), new(table.Servers), new(table.Services))
	base.InitDB(driverName, dataSourceUrl, useCache)
	TableInit()
	UserInit()
}

func TableInit() {
	log.Println("TableInit")
	err := base.DBEngine.Sync2(new(table.User),
		new(table.Server),
		new(table.Service),
		new(table.ServiceReplicas),
		new(table.Container),
		new(table.ContainerStats),
		new(table.Image),
		new(table.Task),
		new(table.Config))
	if err != nil {
		log.Println("TableInit.err:", err)
	}

	log.Println("TableInit Success!")
}
