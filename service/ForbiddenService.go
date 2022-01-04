package service

import (
	"docker-manager/model"
	"log"
)

var Log model.SyncMap

func AddForbiddenLog(reqIp string) {
	log.Println("禁止访问,IP:", reqIp)
	Log.IncInt(reqIp, 1)
}

func ForbiddenLogMap() map[string]interface{} {
	return Log.ToStrMap()
}
