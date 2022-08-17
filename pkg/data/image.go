package data

import (
	"docker-manager/pkg/data/base"
	"docker-manager/pkg/data/table"
	"docker-manager/pkg/model"
	"log"
)

func AddImage(e table.Image) (err error) {
	var record table.Image
	has, err := base.DBEngine.Table("image").Where("server_name=? and image_id=?", e.ServerName, e.ImageId).Get(&record)
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

func DelImage(e table.Image) (err error) {
	affected, err := base.DBEngine.Table("image").Delete(&e)
	if err != nil {
		log.Println("DelImage.err:", err)
		return err
	}
	if affected == 0 {
		log.Println("DelImage.affected is 0")
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

func GetImages(serverNames []string, tagName string, page model.Page) (record []table.Image, tatol int64, err error) {
	session := base.DBEngine.Table("image")
	if len(serverNames) > 0 {
		base.SetArrayParams(session, "server_name", serverNames)
	}

	if tagName != "" {
		session.Where("image_id=?", tagName)
	}
	page.SetPageSql(session)

	tatol, err = session.FindAndCount(&record)
	if err != nil {
		log.Println("GetImages.FindAndCount:", err)
	}
	return
}

func GetImage(imageId string) (record table.Image, err error) {
	_, err = base.DBEngine.Table("image").Where("image_id=?", imageId).Get(&record)
	if err != nil {
		log.Println("AddTask.err:", err)
	}
	return
}
