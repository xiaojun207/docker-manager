package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	"log"
)

func AddReplicas(serviceName, serverName string) (err error) {
	e := table.ServiceReplicas{
		ServiceName: serviceName,
		ServerName:  serverName,
	}
	var record table.ServiceReplicas
	has, err := base.DBEngine.Table("service_replicas").Where("service_name=? and server_name=?", serviceName, serverName).Get(&record)
	if err != nil {
		log.Println("AddReplicas.err:", err)
		return err
	}
	if has {
		base.DBEngine.Table("service_replicas").ID(record.Id).Update(e)
	} else {
		_, err = base.DBEngine.Table("service_replicas").Insert(&e)
	}
	return
}

func DeleteReplicas(id int) (err error) {
	log.Println("DeleteReplicas.id:", id)
	_, err = base.DBEngine.Exec("delete from service_replicas where id=?", id)
	base.DBEngine.ClearCache(new(table.ServiceReplicas))
	return
}

func GetServiceReplicasSize() int64 {
	count, err := base.DBEngine.Table("service_replicas").Count()
	if err != nil {
		log.Println("GetTaskSize.err:", err)
		return 0
	}
	return count
}

func GetServiceReplicas() (record []table.ServiceReplicas, err error) {
	err = base.DBEngine.Table("service_replicas").Find(&record)
	return
}
