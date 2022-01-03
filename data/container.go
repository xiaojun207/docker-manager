package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func AddContainer(e *table.Container) (err error) {
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
		_, err = base.DBEngine.Table("container").Insert(e)
	}
	return
}

func DelContainer(e table.Container) (err error) {
	affected, err := base.DBEngine.Table("container").Delete(&e)
	if err != nil {
		log.Println("DelContainer.err:", err)
		return err
	}
	if affected == 0 {
		log.Println("DelContainer.affected is 0")
	}
	return
}

func GetContainer(ContainerId string) (record table.Container, err error) {
	_, err = base.DBEngine.Table("container").Where("container_id=?", ContainerId).Get(&record)
	return
}

func GetContainers() (record []table.Container, err error) {
	err = base.DBEngine.Table("container").OrderBy("server_name asc,name asc").Find(&record)
	return
}

func GetContainersByServerName(serverName string) (record []table.Container, err error) {
	err = base.DBEngine.Table("container").Where("server_name=?", serverName).OrderBy("name asc").Find(&record)
	return
}

func GetServerNameByContainerShortId(cId string) string {
	var record table.Container
	has, err := base.DBEngine.Table("container").Where("substr(container_id,1,12)=?", cId).Get(&record)
	if err != nil {
		log.Println("GetServerNameByContainerShortId.err:", err)
		return ""
	}
	if has {
		return record.ServerName
	}
	return ""
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
	follow_if := "if(follow=1, 1, null)"
	if base.IsSqlite3() {
		follow_if = "CASE follow WHEN 1 THEN 1 END"
	}
	state_if := "if(state='running', 1, null)"
	if base.IsSqlite3() {
		state_if = "CASE state WHEN 'running' THEN 1 END"
	}
	sql := "select " +
		"count(" + follow_if + ") followSize, " +
		"count(*) totalSize," +
		"count(" + state_if + ") runningSize" +
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
