package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	"errors"
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
)

func FindUsers() (users []table.User, err error) {
	err = base.DBEngine.Table("user").Find(&users)
	if err != nil {
		log.Println("FindUsers.err:", err)
	}
	return
}

func FindUserByUsername(username string) (user table.User, err error) {
	has, err := base.DBEngine.Table("user").Where("username=?", username).Get(&user)
	if err != nil {
		log.Println("FindUserByUsername.err:", err)
		return user, err
	}
	if !has {
		return user, errors.New("FindUserByUsername user '" + username + "' not exists")
	}
	return
}

func UpdatePasswordByUsername(user table.User) (err error) {
	_, err = base.DBEngine.Exec("update user set password=? where username=?", user.Password, user.Username)
	if err != nil {
		log.Println("UpdatePasswordByUsername.err:", err)
		return err
	}
	return
}

func CreateUser(user table.User) error {
	_, err := base.DBEngine.Table("user").Insert(&user)
	if err != nil {
		log.Println("insert user err:", err)
	}
	return err
}

func UserInit() {
	var record table.User
	has, err := base.DBEngine.Table("user").Where("1=1").Limit(1).Get(&record)
	if err != nil {
		log.Println("InitUser.err:", err)
	}
	if has {
		log.Println("User is exists")
	} else {
		adminPassword := utils.Md5(uuid.New())
		adminUser := table.User{
			Username: "admin",
			Slat:     utils.Sha256(uuid.New()),
			Role:     table.USERROLE_ADMIN,
			Nickname: "超级管理员",
		}
		adminUser.Password = adminUser.CreatePassword(adminPassword)
		_, err = base.DBEngine.Table("user").Insert(&adminUser)
		if err != nil {
			log.Println("InitUser insert admin err:", err)
		} else {
			log.Println("InitUser insert [username:", adminUser.Username, ", password:", adminPassword, "], This password only show once!")
		}

		agentPassword := utils.Md5(uuid.New())
		agentUser := table.User{
			Username: "agent",
			Slat:     utils.Sha256(uuid.New()),
			Role:     table.USERROLE_AGENT,
			Nickname: "docker agent",
		}
		agentUser.Password = agentUser.CreatePassword(agentPassword)
		_, err = base.DBEngine.Table("user").Insert(&agentUser)
		if err != nil {
			log.Println("InitUser insert agent err:", err)
		} else {
			log.Println("InitUser insert [username:", agentUser.Username, ", password:", agentPassword, "], This password only show once!")
		}
	}
	return
}
