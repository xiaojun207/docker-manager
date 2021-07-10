package table

import (
	"time"
)

type User struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	Username   string    `xorm:"comment('用户名') VARCHAR(40)"`
	Mobile     string    `xorm:"default '' comment('电话') VARCHAR(20)"`
	Email      string    `xorm:"comment('邮箱') VARCHAR(64)"`
	Password   string    `xorm:"comment('密码') VARCHAR(128)"`
	Slat       string    `xorm:"default '' comment('密码加密盐') VARCHAR(128)"`
	Nickname   string    `xorm:"comment('昵称') VARCHAR(40)"`
	Status     int       `xorm:"default 0 comment('状态，0 正常，1 禁用') INT"`
	Role       string    `xorm:"default 'USER' comment('角色，ADMIN 超级管理员，USER 普通用户，AGENT 代理客户端，DEPLOY 发布api') ENUM('ADMIN','AGENT','DEPLOY','USER')"`
	Summary    string    `xorm:"JSON"`
	CreateDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
