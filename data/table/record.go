package table

import "time"

type Record struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	Uid        int       `xorm:"default 1 comment('uid') INT"`
	IP         string    `xorm:"comment('ip') VARCHAR(24)"`
	Url        string    `xorm:"comment('url') VARCHAR(64)"`
	CreateDate time.Time `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
}
