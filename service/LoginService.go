package service

import (
	"docker-manager/data"
	"docker-manager/model"
	"docker-manager/utils"
	"errors"
)

var userLock = model.NewUserLock()

/**
this is a sample login
*/
func Login(username, password string) (string, error) {
	if userLock.Locked(username) {
		return "", errors.New("用户已被锁定")
	}
	pass, has := data.SampleUser.LoadStr(username)
	if has && password == pass {
		token := utils.CreateToken(username)
		return token, nil
	} else {
		userLock.AddLockNum(username)
	}
	return "", errors.New("用户密码错误")
}
