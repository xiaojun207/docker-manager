package data

import (
	"docker-manager/pkg/data/base"
	"docker-manager/pkg/data/table"
	"docker-manager/pkg/model"
	"log"
	"strconv"
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

func InsertContainerStats(e *table.ContainerStats) (err error) {
	_, err = base.DBEngine.Table("container_stats").Insert(e)
	return
}

func UpdateContainerStats(e table.ContainerStats) (err error) {
	_, err = base.DBEngine.Table("container_stats").ID(e.Id).Update(e)
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

func GetContainerStatsList(Follow string, serverNames, containerNames []string, page *model.Page) (record []table.ContainerStats, err error) {
	session := base.DBEngine.Table("container_stats")
	if Follow != "" {
		f, _ := strconv.ParseBool(Follow)
		session.Where("follow=?", f)
	}
	if len(serverNames) > 0 {
		base.SetArrayParams(session, "server_name", serverNames)
	}
	if len(containerNames) > 0 {
		base.SetArrayParams(session, "name", containerNames)
	}

	err = page.FindPage(session, &record)
	return
}

func DeleteStats(id int) (err error) {
	_, err = base.DBEngine.Table("container_stats").Exec("delete from container_stats where id=?", id)
	if err != nil {
		log.Println("DelStats.err:", err)
	}
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
