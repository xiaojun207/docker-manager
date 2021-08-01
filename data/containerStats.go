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

func GetContainerStats() (record []table.ContainerStats, err error) {
	err = base.DBEngine.Table("container_stats").Find(&record)
	return
}
