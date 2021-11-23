package table

import (
	"docker-manager/model"
	"github.com/xiaojun207/go-base-utils/utils"
)

const (
	USERROLE_ADMIN       = "ADMIN"
	USERROLE_AGENT       = "AGENT"
	USERROLE_DEPLOY      = "DEPLOY"
	USERROLE_USER        = "USER"
	USERSTATUS_NORMAL    = 0
	USERSTATUS_FORBIDDEN = 1
)

type User struct {
	Id         int                    `xorm:"not null pk autoincr INT" json:"id"`
	Username   string                 `xorm:"comment('用户名') VARCHAR(40) unique" json:"username"`
	Mobile     string                 `xorm:"default '' comment('电话') VARCHAR(20)" json:"mobile"`
	Email      string                 `xorm:"comment('邮箱') VARCHAR(64)" json:"email"`
	Password   string                 `xorm:"comment('密码') VARCHAR(128)" json:-`
	Slat       string                 `xorm:"default '' comment('密码加密盐') VARCHAR(128)" json:-`
	Nickname   string                 `xorm:"comment('昵称') VARCHAR(40)" json:"nickname"`
	Status     int                    `xorm:"default 0 comment('状态，0 正常，1 禁用') INT" json:"status"`
	Role       string                 `xorm:"default 'USER' comment('角色，ADMIN 超级管理员，USER 普通用户，AGENT 代理客户端，DEPLOY 发布api') ENUM('ADMIN','AGENT','DEPLOY','USER')" json:"role"`
	Summary    map[string]interface{} `xorm:"JSON" json:"summary"`
	CreateDate model.Time             `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP" json:"create_date"`
	UpdateDate model.Time             `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP" json:"update_date"`
}

func (e User) CreatePassword(inputPassword string) string {
	return utils.Sha256(inputPassword + e.Slat)
}

func (e User) CheckPassword(inputPassword string) bool {
	return e.CreatePassword(inputPassword) == e.Password
}
