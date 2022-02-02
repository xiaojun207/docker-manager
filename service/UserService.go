package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/model"
	sutils "docker-manager/utils"
	"errors"
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
	"regexp"
)

func InitTokenHelper() {
	tokenSecret := data.GetConfigValue("server.tokenSecret", uuid.New(), true)
	sutils.InitTokenHelper(tokenSecret)
}

func FindUsers(page *model.Page) (users []table.User, err error) {
	return data.FindUsers(page)
}

func FindUserByUsername(username string) (user table.User, err error) {
	return data.FindUserByUsername(username)
}

func FindUser(uid int) (user table.User, err error) {
	return data.FindUserByUid(uid)
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func AddUser(nickname, username, mobile, email, password, role string) error {
	if username == "" && mobile == "" && email == "" {
		return errors.New("username, mobile, email 不能同时为空")
	}

	if username != "" {
		if len(username) < 3 || len(username) > 20 {
			return errors.New("username长度4-20")
		}
	}

	if email != "" {
		if !VerifyEmailFormat(email) {
			errors.New("email格式不正确")
		}
	}

	if mobile != "" {
		if len(mobile) < 6 || len(mobile) > 18 {
			return errors.New("mobile长度不正确")
		}
	}

	if len(password) < 6 {
		return errors.New("密码太短")
	}

	if role != table.USERROLE_ADMIN && role != table.USERROLE_AGENT {
		return errors.New("角色错误")
	}

	tmp, err := data.FindUserByUsername(username)
	if tmp.Id > 0 {
		log.Println("AddUser.err username 已存在:", username, ",id:", tmp.Id)
		return errors.New(username + "用户已存在")
	}

	tmp, err = data.FindUserByMobile(mobile)
	if tmp.Id > 0 {
		log.Println("AddUser.err mobile 已存在:", mobile, ",id:", tmp.Id)
		return errors.New(mobile + "用户已存在")
	}

	tmp, err = data.FindUserByEmail(email)
	if tmp.Id > 0 {
		log.Println("AddUser.err email 已存在:", email, ",id:", tmp.Id)
		return errors.New(email + "用户已存在")
	}

	user := table.User{
		Username: username,
		Mobile:   mobile,
		Email:    email,
		Slat:     utils.Sha256(uuid.New()),
		Role:     role,
		Nickname: nickname,
	}
	user.Password = user.CreatePassword(password)
	err = data.CreateUser(user)
	if err != nil {
		return errors.New("新增用户异常")
	}
	return nil
}

func AlterPassword(uid int, oldPassword, newPassword string) error {
	if len(newPassword) < 6 {
		return errors.New("密码太短")
	}

	user, err := data.FindUserByUid(uid)
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
