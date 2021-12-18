package table

import (
	"time"
)

type Server struct {
	Id              int    `xorm:"not null pk autoincr INT"`
	Name            string `xorm:"comment('服务器名称') VARCHAR(40) unique"`
	OSType          string `xorm:"comment('操作系统类型') VARCHAR(40)"`
	OperatingSystem string `xorm:"comment('操作系统') VARCHAR(40)"`
	KernelVersion   string `xorm:"comment('操作系统核心版本') VARCHAR(40)"`
	DockerVersion   string `xorm:"comment('docker版本') VARCHAR(40)"`
	Running         int    `xorm:"default 0 comment('运行中的容器') INT"`
	Containers      int    `xorm:"default 0 comment('容器数量') INT"`
	Images          int    `xorm:"default 0 comment('镜像数量') INT"`
	Cpu             int    `xorm:"default 1 comment('cpu核心线程数量') INT"`
	Memory          int64  `xorm:"default 1024 comment('内存') BIGINT"`
	PrivateIp       string `xorm:"comment('内网ip') VARCHAR(24)"`
	PublicIp        string `xorm:"comment('公网ip') VARCHAR(24)"`
	Status          string `xorm:"comment('状态1， 不在线，在线') VARCHAR(20)"`
	State           string `xorm:"comment('状态2， 不在线，在线') VARCHAR(20)"`
	LastDataTime    int64
	Summary         map[string]interface{} `xorm:"JSON"`
	CreateDate      time.Time              `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate      time.Time              `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}
