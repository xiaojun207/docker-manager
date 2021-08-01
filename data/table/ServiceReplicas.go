package table

import (
	"time"
)

type ServiceReplicas struct {
	Id          int       `xorm:"not null pk autoincr INT"`
	ServiceName string    `xorm:"comment('服务名称（同容器名称）') VARCHAR(40) unique"`
	ServerName  string    `xorm:"comment('镜像名称') VARCHAR(1024)"`
	Status      int       `xorm:"default 0 comment('状态，0 不在线，1在线') INT"`
	Summary     string    `xorm:"JSON"`
	CreateDate  time.Time `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate  time.Time `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}
