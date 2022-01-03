package conf

import (
	"flag"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

var (
	DriverName    = "sqlite3"
	DataSourceUrl = "data/database.db"
	UseCache      = "true"
	ConsoleCode   = "false"
)

func ParseParam() {
	flag.StringVar(&DriverName, "driverName", DriverName, "driverName, eg.: mysql or sqlite3")
	flag.StringVar(&UseCache, "useCache", UseCache, "useCache, eg.: true or false")
	flag.StringVar(&ConsoleCode, "consoleCode", ConsoleCode, "consoleCode, eg.: true or false")
	flag.StringVar(&DataSourceUrl, "dataSourceUrl", DataSourceUrl, "dataSourceUrl, eg.: root:123456@(127.0.0.1:3306)/docker-manager?charset=utf8")
	flag.Parse()
	log.Println("flag.driverName:", DriverName)
	log.Println("flag.useCache:", UseCache)
	log.Println("flag.consoleCode:", ConsoleCode)
	log.Println("flag.dataSourceUrl:", utils.SubstrAfter(DataSourceUrl, "@"))
}
