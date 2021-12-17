package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	"log"
)

func AddContainerStats(e table.ContainerStats) (err error) {
	var record table.ContainerStats
	has, err := base.DBEngine.Table("container_stats").Where("container_id=?", e.ContainerId).Get(&record)
	if err != nil {
		log.Println("AddContainerStats.err:", err)
		return err
	}
	if has {
		base.DBEngine.Table("container_stats").ID(record.Id).Update(e)
	} else {
		_, err = base.DBEngine.Table("container_stats").Insert(&e)
	}
	return
}

func GetContainerStatsSize() int64 {
	count, err := base.DBEngine.Table("container_stats").Count()
	if err != nil {
		log.Println("GetContainerStatsSize.err:", err)
		return 0
	}
	return count
}

func GetContainerStatss(ContainerId string) (record table.ContainerStats, err error) {
	_, err = base.DBEngine.Table("container_stats").Where("container_id=?", ContainerId).Get(&record)
	return
}

func GetContainerStatsList(serverNames []string) (record []table.ContainerStats, err error) {
	session := base.DBEngine.Table("container_stats")
	if len(serverNames) > 0 {
		sql, params := base.ArrayParams(serverNames)
		session.Where("server_name in ("+sql+")", params...)
	}

	err = session.Find(&record)
	return
}

func DelStats(e table.ContainerStats) (err error) {
	affected, err := base.DBEngine.Table("container_stats").Delete(&e)
	if err != nil {
		log.Println("DelStats.err:", err)
		return err
	}
	if affected == 0 {
		log.Println("DelStats.affected is 0")
	}
	return
}
