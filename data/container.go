package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func AddContainer(e table.Container) (err error) {
	var record table.Container
	has, err := base.DBEngine.Table("container").
		Where("container_id=? or (name=? and server_name=?)", e.ContainerId, e.Name, e.ServerName).
		Get(&record)
	if err != nil {
		log.Println("AddContainer.err:", err)
		return err
	}
	if has {
		_, err = base.DBEngine.Table("container").ID(record.Id).Update(e)
	} else {
		_, err = base.DBEngine.Table("container").Insert(&e)
	}
	return
}

func GetContainers() (record []table.Container, err error) {
	err = base.DBEngine.Table("container").OrderBy("server_name asc,name asc").Find(&record)
	return
}

func GetServerNameByContainerId(containerId string) string {
	var record table.Container
	has, err := base.DBEngine.Table("container").Where("container_id=?", containerId).Get(&record)
	if err != nil {
		log.Println("GetServerNameByContainerId.err:", err)
		return ""
	}
	if has {
		return record.ServerName
	}
	return ""
}

func GetContainerSize() map[string]int64 {
	sql := "select " +
		"count(if(follow=1, 1, null)) followSize, " +
		"count(*) totalSize," +
		"count(if(state='running', 1, null)) runningSize" +
		" from container"
	resMap, err := base.DBEngine.QueryString(sql)
	if err != nil {
		log.Println("GetContainerSize.err:", err)
		return nil
	}
	s := resMap[0]
	return map[string]int64{
		"followSize":  utils2.StrToInt64(s["followSize"]),
		"totalSize":   utils2.StrToInt64(s["totalSize"]),
		"runningSize": utils2.StrToInt64(s["runningSize"]),
	}
}
