package table

import "time"

type ContainerStats struct {
	Id           int       `xorm:"not null pk autoincr INT"`
	ContainerId  string    `xorm:"comment('容器Id') VARCHAR(140)"`
	Name         string    `xorm:"comment('容器名称') VARCHAR(40)"`
	ServerName   string    `xorm:"comment('所在服务器名称') VARCHAR(40)"`
	Time         int64     `xorm:"default 0 comment('time') INT"`
	BlkioStats   string    `xorm:"JSON"`
	CpuStats     string    `xorm:"JSON"`
	MemoryStats  string    `xorm:"JSON"`
	Networks     string    `xorm:"JSON"`
	NumProcs     int64     `xorm:"default 0 comment('Num Procs') INT"`
	PidsStats    string    `xorm:"JSON"`
	PrecpuStats  string    `xorm:"JSON"`
	Preread      string    `xorm:"comment('Preread') VARCHAR(128)"`
	Read         string    `xorm:"comment('read') VARCHAR(128)"`
	StorageStats string    `xorm:"JSON"`
	Follow       bool      `xorm:"default 0 comment('是否日志跟随') Bool"`
	Summary      string    `xorm:"JSON"`
	CreateDate   time.Time `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate   time.Time `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}
