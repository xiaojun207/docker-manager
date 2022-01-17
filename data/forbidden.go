package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
)

func AddForbidden(e table.Forbidden) (err error) {
	_, err = base.DBEngine.Table("forbidden").Insert(&e)
	return
}

func ForbiddenList() (records []table.Forbidden) {
	base.DBEngine.Table("forbidden").SQL("select ip,sum(num) as num,max(create_date) as create_date from forbidden").Find(&records)
	return
}
