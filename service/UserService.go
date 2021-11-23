package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"errors"
	"github.com/xiaojun207/go-base-utils/utils"
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

func ResetPassword(username string) (string, error) {
	user, err := data.FindUserByUsername(username)
	if err != nil {
		log.Println("ResetPassword.err:", err)
		return "", errors.New("用户不存在")
	}

	newPassword := utils.RandomPassword(32, "mix")
	user.Password = user.CreatePassword(newPassword)
	return newPassword, data.UpdatePasswordByUsername(user)

}

func ChangeStatus(username string, status int) error {
	user, err := data.FindUserByUsername(username)
	if err != nil {
		log.Println("ChangeStatus.err:", err)
		return errors.New("用户不存在")
	}

	user.Status = status
	return data.UpdateStatusByUsername(user)

}
