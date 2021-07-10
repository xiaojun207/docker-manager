package table

import (
	"time"
)

type UserApi struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	Uid        int       `xorm:"not null comment('user id') INT"`
	Name       string    `xorm:"default '' comment('api名称') VARCHAR(40)"`
	ApiKey     string    `xorm:"default '' comment('api key') VARCHAR(40)"`
	ApiSecret  string    `xorm:"default '' comment('api秘钥') VARCHAR(40)"`
	Status     int       `xorm:"default 0 comment('状态，0 正常，1 禁用') INT"`
	Summary    string    `xorm:"JSON"`
	CreateDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
