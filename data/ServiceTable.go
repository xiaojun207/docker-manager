package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	"docker-manager/dto"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func AddAppInfo(appname string, info dto.ServiceInfo) {
	//if AppInfos.Has(appname) {
	//	//return
	//}
	//AppInfos.Store(appname, info)

	service := table.Service{
		Name:  appname,
		Image: info.Image,
		//Ports:    utils2.StructToMap(info.Ports),
		//Env:      utils2.StructToJson(info.Env),
		//Vol:      utils2.StructToJson(info.Volumes),
		Running:  info.Running,
		Replicas: info.Replicas,
		//Status: info.
		//Summary: info.
	}
	utils2.StructToMap(info.Ports, &service.Ports)
	utils2.StructToMap(info.Env, &service.Env)
	utils2.StructToMap(info.Volumes, &service.Vol)
	AddService(service)
}

func AddService(service table.Service) (err error) {
	var record table.Service
	has, err := base.DBEngine.Table("service").Where("Name=?", service.Name).Get(&record)
	if err != nil {
		log.Println("AddService.err:", err)
		return err
	}
	if has {
		base.DBEngine.Table("service").ID(record.Id).Update(service)
	} else {
		_, err = base.DBEngine.Table("service").Insert(&service)
	}
	return
}

func DeleteService(serviceName string) (err error) {
	_, err = base.DBEngine.Exec("delete from service where Name=?", serviceName)
	base.DBEngine.ClearCache(new(table.Service))
	return
}

func GetServices() (record []table.Service, err error) {
	err = base.DBEngine.Table("service").Find(&record)
	return
}

func GetService(name string) (record table.Service, err error) {
	_, err = base.DBEngine.Table("service").Where("Name=?", name).Get(&record)
	if err != nil {
		log.Println("AddService.err:", err)
	}
	return
}
func GetServiceSize() int64 {
	count, err := base.DBEngine.Table("service").Count()
	if err != nil {
		log.Println("GetServiceSize.err:", err)
		return 0
	}
	return count
}
