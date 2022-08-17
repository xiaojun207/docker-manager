package table

import (
	"time"
)

type Config struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	Name       string    `xorm:"comment('字段名称') VARCHAR(64) unique"`
	Value      string    `xorm:"comment('值') JSON"`
	Memo       string    `xorm:"comment('备注') VARCHAR(64)"`
	CreateDate time.Time `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate time.Time `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}
