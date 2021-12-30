package service

import (
	"docker-manager/model"
	"github.com/xiaojun207/go-base-utils/utils"
)

var codeCacheMap = model.SyncMap{}

func SendCode(uid int, codeType, key string) string {
	code := utils.RandomPassword(6, "mix")
	codeCacheMap.Store(uid, code)
	return code
}

func CheckCode(uid int, code string) bool {
	cacheCode, has := codeCacheMap.LoadAndDelete(uid)
	return has && cacheCode == code
}
