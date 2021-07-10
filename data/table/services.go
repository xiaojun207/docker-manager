package table

import (
	"time"
)

type Services struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	Name       string    `xorm:"comment('服务名称（同容器名称）') VARCHAR(40)"`
	Image      string    `xorm:"comment('镜像名称') VARCHAR(1024)"`
	Ports      string    `xorm:"comment('端口列表') JSON"`
	Env        string    `xorm:"comment('环境变量列表') JSON"`
	Vol        string    `xorm:"comment('路径映射列表') JSON"`
	Running    int       `xorm:"default 0 comment('运行实例数量') INT"`
	Replicas   int       `xorm:"default 0 comment('副本数量') INT"`
	Status     int       `xorm:"default 0 comment('状态，0 不在线，1在线') INT"`
	Summary    string    `xorm:"JSON"`
	CreateDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
