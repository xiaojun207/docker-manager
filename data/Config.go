package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func GetConfigValue(name, defValue string) string {
	var conf table.Config
	has, err := base.DBEngine.Table("config").Where("name=?", name).Get(&conf)
	if err != nil {
		log.Println("GetConfigValue.err:", err)
		return defValue
	}
	if !has {
		return defValue
	}
	return conf.Value
}

func GetAgentConfig() map[string]interface{} {
	conf := map[string]interface{}{
		"TaskFrequency": utils2.StrToFloat64(GetConfigValue("agent.TaskFrequency", "300")),
	}
	return map[string]interface{}{
		"agentConfig": conf,
	}
}

func GetConfigList() (record []table.Config) {
	base.DBEngine.Table("config").Find(&record)
	return
}

func UpdateConfig(name, value, memo string) (err error) {
	var record table.Config
	has, err := base.DBEngine.Table("config").Where("name=?", name).Get(&record)
	if err != nil {
		log.Println("UpdateConfig.err:", err)
		return err
	}
	if has {
		record.Name = name
		record.Value = value
		record.Memo = memo
		base.DBEngine.Table("task").ID(record.Id).Update(record)
	}
	return
}
