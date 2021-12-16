package table

import (
	"docker-manager/model"
)

type ContainerStats struct {
	Id           int                    `xorm:"not null pk autoincr INT"`
	ContainerId  string                 `xorm:"comment('容器Id') VARCHAR(140)"`
	Name         string                 `xorm:"comment('容器名称') VARCHAR(40)"`
	ServerName   string                 `xorm:"comment('所在服务器名称') VARCHAR(40)"`
	Time         int64                  `xorm:"default 0 comment('time') INT"`
	BlkioStats   map[string]interface{} `xorm:"JSON" json:"blkio_stats"`
	CpuStats     map[string]interface{} `xorm:"JSON" json:"cpu_stats"`
	MemoryStats  map[string]interface{} `xorm:"JSON" json:"memory_stats"`
	Networks     map[string]interface{} `xorm:"JSON" json:"networks"`
	NumProcs     int64                  `xorm:"default 0 comment('Num Procs') INT"`
	PidsStats    map[string]interface{} `xorm:"JSON" json:"pids_stats"`
	PrecpuStats  map[string]interface{} `xorm:"JSON" json:"precpu_stats"`
	Preread      string                 `xorm:"comment('Preread') VARCHAR(128)"`
	Read         string                 `xorm:"comment('read') VARCHAR(128)"`
	StorageStats map[string]interface{} `xorm:"JSON" json:"storage_stats"`
	Follow       bool                   `xorm:"default 0 comment('是否日志跟随') Bool"`
	Summary      map[string]interface{} `xorm:"JSON"`
	CreateDate   model.Time             `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate   model.Time             `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}
