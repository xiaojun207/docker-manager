package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	"docker-manager/model"
	"log"
)

func AddTask(e table.Task) (err error) {
	var record table.Task
	has, err := base.DBEngine.Table("task").Where("task_id=?", e.TaskId).Get(&record)
	if err != nil {
		log.Println("AddTask.err:", err)
		return err
	}
	if has {
		base.DBEngine.Table("task").ID(record.Id).Update(e)
	} else {
		_, err = base.DBEngine.Table("task").Insert(&e)
	}
	return
}

func UpdateTask(taskId, code, msg string, resp map[string]interface{}) (err error) {
	var record table.Task
	has, err := base.DBEngine.Table("task").Where("task_id=?", taskId).Get(&record)
	if err != nil {
		log.Println("UpdateTask.err:", err)
		return err
	}
	if has {
		record.Code = code
		record.Msg = msg
		record.Resp = resp
		base.DBEngine.Table("task").ID(record.Id).Update(record)
	}
	return
}

func GetTaskSize() int64 {
	count, err := base.DBEngine.Table("task").Count()
	if err != nil {
		log.Println("GetTaskSize.err:", err)
		return 0
	}
	return count
}

func GetTasks(page *model.Page) (record []table.Task, err error) {
	session := base.DBEngine.Table("task").OrderBy("id desc")
	page.SetPageSql(session)
	page.Total, err = session.FindAndCount(&record)
	if err != nil {
		log.Println("GetTasks.FindAndCount:", err)
	}
	return
}
