package table

import (
	"github.com/xiaojun207/go-base-utils/utils"
	"time"
)

const (
	USERROLE_ADMIN  = "ADMIN"
	USERROLE_AGENT  = "AGENT"
	USERROLE_DEPLOY = "DEPLOY"
	USERROLE_USER   = "USER"
)

type User struct {
	Id         int                    `xorm:"not null pk autoincr INT"`
	Username   string                 `xorm:"comment('用户名') VARCHAR(40) unique"`
	Mobile     string                 `xorm:"default '' comment('电话') VARCHAR(20)"`
	Email      string                 `xorm:"comment('邮箱') VARCHAR(64)"`
	Password   string                 `xorm:"comment('密码') VARCHAR(128)"`
	Slat       string                 `xorm:"default '' comment('密码加密盐') VARCHAR(128)"`
	Nickname   string                 `xorm:"comment('昵称') VARCHAR(40)"`
	Status     int                    `xorm:"default 0 comment('状态，0 正常，1 禁用') INT"`
	Role       string                 `xorm:"default 'USER' comment('角色，ADMIN 超级管理员，USER 普通用户，AGENT 代理客户端，DEPLOY 发布api') ENUM('ADMIN','AGENT','DEPLOY','USER')"`
	Summary    map[string]interface{} `xorm:"JSON"`
	CreateDate time.Time              `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate time.Time              `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (e User) CreatePassword(inputPassword string) string {
	return utils.Sha256(inputPassword + e.Slat)
}

func (e User) CheckPassword(inputPassword string) bool {
	return e.CreatePassword(inputPassword) == e.Password
}
