package service

import (
	"docker-manager/pkg/model"
	"github.com/xiaojun207/go-base-utils/utils"
	"time"
)

// only a little data
var codeCacheMap = model.SyncMap{}

type ValCode struct {
	code string
	key  string
	t    time.Time
}

func SendCode(uid int, codeType, key string) string {
	code := utils.RandomPassword(6, "mix")
	varCode := ValCode{
		code: code,
		key:  key,
		t:    time.Now(),
	}
	codeCacheMap.Store(uid, varCode)
	return code
}

func CheckCode(uid int, code string) bool {
	cacheCode, has := codeCacheMap.LoadAndDelete(uid)
	if has {
		varCode := cacheCode.(ValCode)
		return varCode.code == code && varCode.t.Add(time.Minute).After(time.Now())
	}
	return false
}
