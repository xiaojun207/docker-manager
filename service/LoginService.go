package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/model"
	"docker-manager/utils"
	"errors"
	"log"
)

var userLock = model.NewUserLock()

/**
this is a sample login
*/
func Login(username, password string) (string, error) {
	if userLock.Locked(username) {
		return "", errors.New("用户已被锁定")
	}

	user, err := data.FindUserByUsername(username)
	if err != nil {
		log.Println("Login.err:", err)
		return "", errors.New("用户密码错误")
	}

	if user.Status != table.USERSTATUS_NORMAL {
		return "", errors.New("用户密码错误!!")
	}

	if user.Role != table.USERROLE_ADMIN && user.Role != table.USERROLE_USER {
		// 角色不匹配
		return "", errors.New("用户密码错误!")
	}

	// 密码检测
	if user.CheckPassword(password) {
		token := utils.CreateToken(username)
		return token, nil
	} else {
		userLock.AddLockNum(username)
	}
	return "", errors.New("用户密码错误")
}

func LoginApi(apiKey, apiSecret string) bool {
	user, err := data.FindUserByUsername(apiKey)
	if err != nil {
		log.Println("LoginApi.err:", err)
		return false
	}

	if user.Status != table.USERSTATUS_NORMAL {
		return false
	}

	if user.Role != table.USERROLE_DEPLOY {
		// 角色不匹配
		return false
	}
	return user.CheckPassword(apiSecret)
}

func LoginAgent(agentPassword string) bool {
	username := "agent"

	user, err := data.FindUserByUsername(username)
	if err != nil {
		log.Println("LoginAgent.err:", err)
		return false
	}

	if user.Status != table.USERSTATUS_NORMAL {
		return false
	}

	if user.Role != table.USERROLE_AGENT {
		// 角色不匹配
		return false
	}
	return user.CheckPassword(agentPassword)
}
