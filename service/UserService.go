package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"errors"
	"log"
)

func FindUsers() (users []table.User, err error) {
	return data.FindUsers()
}

func AlterPassword(username, oldPassword, newPassword string) error {
	user, err := data.FindUserByUsername(username)
	if err != nil {
		log.Println("AlterPassword.err:", err)
		return errors.New("用户不存在")
	}

	// 密码检测
	if user.CheckPassword(oldPassword) {
		user.Password = user.CreatePassword(newPassword)
		return data.UpdatePasswordByUsername(user)
	} else {
		return errors.New("旧密码错误")
	}
}
