package table

import (
	"time"
)

type Servers struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	Name       string    `xorm:"comment('服务器名称') VARCHAR(40)"`
	Running    int       `xorm:"default 0 comment('运行中的容器') INT"`
	Containers int       `xorm:"default 0 comment('容器数量') INT"`
	Cpu        int       `xorm:"default 1 comment('cpu核心线程数量') INT"`
	Memory     int64     `xorm:"default 1024 comment('内存') BIGINT"`
	PrivateIp  string    `xorm:"comment('内网ip') VARCHAR(24)"`
	PublicIp   string    `xorm:"comment('公网ip') VARCHAR(24)"`
	Status     int       `xorm:"default 0 comment('状态，0 不在线，1在线') INT"`
	Summary    string    `xorm:"JSON"`
	CreateDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
