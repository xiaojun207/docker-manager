package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"log"
)

func AddForbiddenLog(reqIp string) {
	log.Println("禁止访问,IP:", reqIp)
	data.AddForbidden(table.Forbidden{
		Ip:  reqIp,
		Num: 1,
	})
}

// 禁用记录
func ForbiddenLogRecord() []table.Forbidden {
	return data.ForbiddenList()
}
