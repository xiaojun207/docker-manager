package table

import "time"

type Forbidden struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	Ip         string    `xorm:"comment('ip') VARCHAR(24)"`
	Num        int       `xorm:"default 1 comment('num') INT"`
	CreateDate time.Time `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
}
