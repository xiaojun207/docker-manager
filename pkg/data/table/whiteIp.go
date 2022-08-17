package table

import (
	"time"
)

type WhiteIp struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	IP         string    `xorm:"'ip' comment('ip') VARCHAR(24)"`
	Memo       string    `xorm:"'memo' comment('ip') VARCHAR(24)"`
	CreateDate time.Time `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
}
