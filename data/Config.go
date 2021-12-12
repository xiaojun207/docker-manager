package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
	"strings"
)

const DesKey = "bdea7bcf"

func GetConfigValue(name, defValue string, encrypt bool) string {
	var conf table.Config
	has, err := base.DBEngine.Table("config").Where("name=?", name).Get(&conf)
	if err != nil {
		log.Println("GetConfigValue.err:", err)
		return defValue
	}
	if !has {
		UpdateConfig(name, defValue, "", encrypt)
		return defValue
	}
	if strings.HasPrefix(conf.Value, "DESEncrypt(") {
		value := strings.TrimPrefix(conf.Value, "DESEncrypt(")
		value = strings.TrimSuffix(value, ")")
		return utils2.DESDecrypt(value, DesKey)
	}
	return conf.Value
}

func GetAgentConfig() map[string]interface{} {
	conf := map[string]interface{}{
		"TaskFrequency": utils2.StrToFloat64(GetConfigValue("agent.TaskFrequency", "60", false)),
	}
	return map[string]interface{}{
		"agentConfig": conf,
	}
}

func GetConfigList() (record []table.Config) {
	base.DBEngine.Table("config").Find(&record)
	return
}

func UpdateConfig(name, value, memo string, encrypt bool) (err error) {
	if encrypt {
		value = "DESEncrypt(" + utils2.DESEncrypt(value, DesKey) + ")"
	}

	log.Println("UpdateConfig.name:", name, ",value:", value, ",memo:", memo)
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
		base.DBEngine.Table("config").ID(record.Id).Update(record)
	} else {
		record.Name = name
		record.Value = value
		record.Memo = memo
		i, err := base.DBEngine.Table("config").Insert(record)

		log.Println("UpdateConfig.err:", i, err)
	}
	return
}
