package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	"docker-manager/model"
	"errors"
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
	"strconv"
)

func FindUsers(page *model.Page) (records []table.User, err error) {
	session := base.DBEngine.Table("user")
	page.SetPageSql(session)
	page.Total, err = session.FindAndCount(&records)
	if err != nil {
		log.Println("FindUsers.FindAndCount:", err)
	}
	return
}

func FindUserByUid(uid int) (user table.User, err error) {
	has, err := base.DBEngine.Table("user").Where("id=?", uid).Get(&user)
	if err != nil {
		log.Println("FindUserByUid.err:", err)
		return user, err
	}
	if !has {
		return user, errors.New("FindUserByUid user '" + strconv.Itoa(uid) + "' not exists")
	}
	return
}

func FindUserByUsernameMobileMail(username, mobile, email string) (users []table.User, err error) {
	user := table.User{
		Username: username,
		Mobile:   mobile,
		Email:    email,
	}
	err = base.DBEngine.Table("user").Find(&users, user)
	if err != nil {
		log.Println("FindUserByUsernameMobileMail.err:", err)
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

func FindUserByMobile(mobile string) (user table.User, err error) {
	has, err := base.DBEngine.Table("user").Where("mobile=?", mobile).Get(&user)
	if err != nil {
		log.Println("FindUserByMobile.err:", err)
		return user, err
	}
	if !has {
		return user, errors.New("FindUserByMobile user '" + mobile + "' not exists")
	}
	return
}

func FindUserByEmail(email string) (user table.User, err error) {
	has, err := base.DBEngine.Table("user").Where("email=?", email).Get(&user)
	if err != nil {
		log.Println("FindUserByEmail.err:", err)
		return user, err
	}
	if !has {
		return user, errors.New("FindUserByEmail user '" + email + "' not exists")
	}
	return
}

func UpdatePasswordByUsername(user table.User) (err error) {
	_, err = base.DBEngine.Exec("update user set password=? where username=?", user.Password, user.Username)
	if err != nil {
		log.Println("UpdatePasswordByUsername.err:", err)
		return err
	}
	base.DBEngine.ClearCache(new(table.User))
	return
}

func UpdateStatusByUsername(user table.User) (err error) {
	_, err = base.DBEngine.Exec("update user set status=? where username=?", user.Status, user.Username)
	if err != nil {
		log.Println("UpdateStatusByUsername.err:", err)
		return err
	}
	base.DBEngine.ClearCache(new(table.User))
	return
}

func CreateUser(user table.User) error {
	_, err := base.DBEngine.Table("user").Insert(&user)
	if err != nil {
		log.Println("insert user err:", err)
	}
	return err
}

func DeleteUser(uid int) error {
	user, err := FindUserByUid(uid)
	if err != nil {
		log.Println("DeleteUser.find user err:", err)
	}

	_, err = base.DBEngine.Table("user").Delete(&user)
	if err != nil {
		log.Println("DeleteUser user err:", err)
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
		adminPassword := utils.RandomPassword(32, "mix")
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

		agentPassword := utils.RandomPassword(32, "mix")
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
