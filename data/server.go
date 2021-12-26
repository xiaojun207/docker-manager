package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	"log"
)

func AddServer(server table.Server) (err error) {
	var record table.Server
	has, err := base.DBEngine.Table("server").Where("Name=?", server.Name).Get(&record)
	if err != nil {
		log.Println("AddServer.err:", err)
		return err
	}
	if has {
		_, err = base.DBEngine.Table("server").ID(record.Id).Update(server)
	} else {
		_, err = base.DBEngine.Table("server").Insert(&server)
	}
	return
}

func UpdateServerState(name string, state string) (err error) {
	var record table.Server
	has, err := base.DBEngine.Table("server").Where("Name=?", name).Get(&record)
	if err != nil {
		log.Println("AddServer.err:", err)
		return err
	}
	if has {
		record.State = state
		_, err = base.DBEngine.Table("server").ID(record.Id).Update(record)
	}
	return
}

func DeleteServer(name string) (err error) {
	var record table.Server
	has, err := base.DBEngine.Table("server").Where("Name=?", name).Get(&record)
	if err != nil {
		log.Println("AddServer.err:", err)
		return err
	}
	if has {
		_, err = base.DBEngine.Table("server").Delete(record)
	}
	return
}

func GetServer(ServerName string) (record table.Server, err error) {
	_, err = base.DBEngine.Table("server").Where("Name=?", ServerName).Get(&record)
	return
}

func GetServers() (record []table.Server, err error) {
	err = base.DBEngine.Table("server").OrderBy("name asc").Find(&record)
	return
}

func GetServersName() (record []string, err error) {
	servers, err := GetServers()
	for _, server := range servers {
		record = append(record, server.Name)
	}
	return
}

func GetServersSize() int64 {
	count, err := base.DBEngine.Table("server").Count()
	if err != nil {
		log.Println("GetServersSize.err:", err)
		return 0
	}
	return count
}
