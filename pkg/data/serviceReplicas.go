package data

import (
	"docker-manager/pkg/data/base"
	"docker-manager/pkg/data/table"
	"docker-manager/pkg/model"
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

func GetServiceReplicas(page *model.Page) (record []table.ServiceReplicas, err error) {
	session := base.DBEngine.Table("service_replicas")
	page.FindPage(session, &record)
	return
}
