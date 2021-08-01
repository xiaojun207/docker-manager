package table

import "time"

type Container struct {
	Id              int       `xorm:"not null pk autoincr INT"`
	AppId           string    `xorm:"comment('AppId') VARCHAR(40)"`
	Name            string    `xorm:"comment('容器名称') VARCHAR(40)"`
	ServerName      string    `xorm:"comment('所在服务器名称') VARCHAR(40)"`
	Command         string    `xorm:"comment('Command') VARCHAR(128)"`
	Created         int64     `xorm:"default 0 comment('容器Created时间') INT"`
	HostConfig      string    `xorm:"comment('容器Host配置') JSON"`
	Mounts          string    `xorm:"comment('容器挂载') JSON"`
	NetworkSettings string    `xorm:"comment('容器网络') JSON"`
	Ports           string    `xorm:"comment('容器端口') JSON"`
	ContainerId     string    `xorm:"comment('容器Id') VARCHAR(140)"`
	Image           string    `xorm:"comment('镜像') VARCHAR(240)"`
	ImageID         string    `xorm:"comment('镜像Id') VARCHAR(140)"`
	State           string    `xorm:"default '' comment('状态1') VARCHAR(24)"`
	Status          string    `xorm:"default '' comment('状态2') VARCHAR(24)"`
	Update          int64     `xorm:"default 0 comment('容器Update时间') INT"`
	Follow          bool      `xorm:"default 0 comment('是否日志跟随') Bool"`
	Summary         string    `xorm:"JSON"`
	CreateDate      time.Time `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate      time.Time `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}
