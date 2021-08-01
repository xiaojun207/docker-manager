package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	"log"
)

func AddImage(e table.Image) (err error) {
	var record table.Image
	has, err := base.DBEngine.Table("image").Where("image=?", e.Image).Get(&record)
	if err != nil {
		log.Println("AddTask.err:", err)
		return err
	}
	if has {
		base.DBEngine.Table("image").ID(record.Id).Update(e)
	} else {
		_, err = base.DBEngine.Table("image").Insert(&e)
	}
	return
}

func GetImageSize() int64 {
	count, err := base.DBEngine.Table("image").Count()
	if err != nil {
		log.Println("GetImageSize.err:", err)
		return 0
	}
	return count
}

func GetImages() (record []table.Image, err error) {
	err = base.DBEngine.Table("image").Find(&record)
	return
}
