package conf

import (
	"flag"
	"github.com/xiaojun207/go-base-utils/utils"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

const YamlFile = "config.yml"

var (
	ApplicationName = "docker-manager"
	Port            = "8068"
	ContextPath     = "/dockerMgrApi/"
	ConsoleCode     = "false"
	DriverName      = "sqlite3"
	DataSourceUrl   = "data/database.db"
	UseCache        = "true"
)

func NewConfig() Config {
	return Config{
		Server: ConfigServer{
			Port:            Port,
			ContextPath:     ContextPath,
			ConsoleCode:     ConsoleCode,
			ApplicationName: ApplicationName,
		},
		DB: ConfigDB{
			DriverName:    DriverName,
			DataSourceUrl: DataSourceUrl,
			UseCache:      UseCache,
		},
	}
}

func ParseParam() {
	c, _ := LoadConfig(YamlFile)

	flag.StringVar(&c.Server.Port, "port", c.Server.Port, "port, default: 8068")
	flag.StringVar(&c.Server.ContextPath, "contextPath", c.Server.ContextPath, "contextPath, default: /dockerMgrApi/")
	flag.StringVar(&c.Server.ConsoleCode, "consoleCode", c.Server.ConsoleCode, "consoleCode, eg.: true or false")
	flag.StringVar(&c.DB.DriverName, "driverName", c.DB.DriverName, "driverName, eg.: mysql or sqlite3")
	flag.StringVar(&c.DB.UseCache, "useCache", c.DB.UseCache, "useCache, eg.: true or false")
	flag.StringVar(&c.DB.DataSourceUrl, "dataSourceUrl", c.DB.DataSourceUrl, "dataSourceUrl, eg.: root:123456@(127.0.0.1:3306)/docker-manager?charset=utf8")
	flag.Parse()
	log.Println("Parameter priority(参数优先级): [HardCode] < [ConfigFile(./config.yml)] < [Cmd parameters (or docker -e)]")

	ApplicationName = c.Server.ApplicationName
	Port = c.Server.Port
	ContextPath = c.Server.ContextPath
	ConsoleCode = c.Server.ConsoleCode
	DriverName = c.DB.DriverName
	DataSourceUrl = c.DB.DataSourceUrl
	UseCache = c.DB.UseCache

	log.Println("flag.ApplicationName:", ApplicationName)
	log.Println("flag.port:", Port)
	log.Println("flag.contextPath:", ContextPath)
	log.Println("flag.consoleCode:", ConsoleCode)
	log.Println("flag.driverName:", DriverName)
	log.Println("flag.dataSourceUrl:", utils.SubstrAfter(DataSourceUrl, "@"))
	log.Println("flag.useCache:", UseCache)

}

func LoadConfig(yamlFile string) (res Config, err error) {
	// 如果异常，则返回默认配置
	res = NewConfig()
	if _, err = os.Stat(yamlFile); err != nil {
		log.Println("configFile is not exits!, File:", yamlFile)
		return res, err
	}

	yamlData, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Println("LoadConfig.err:", err)
		return
	}
	//log.Println("yamlData:\r", string(yamlData))
	err = yaml.Unmarshal(yamlData, &res)
	if err != nil {
		log.Println("LoadConfig.Unmarshal.err:", err)
	}
	return res, err
}
