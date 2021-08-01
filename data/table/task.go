package table

import "time"

type Task struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	TaskId     string    `xorm:"comment('任务id') VARCHAR(40)"`
	Name       string    `xorm:"comment('任务名称') VARCHAR(40)"`
	Ch         string    `xorm:"comment('任务类型') VARCHAR(40)"`
	Code       string    `xorm:"comment('任务返回码') VARCHAR(40)"`
	Msg        string    `xorm:"comment('任务信息') VARCHAR(40)"`
	ServerName string    `xorm:"comment('目标服务器') VARCHAR(40)"`
	Ts         int64     `xorm:"default 0 comment('任务时间') BIGINT"`
	Param      string    `xorm:"comment('任务参数') JSON"`
	Resp       string    `xorm:"comment('任务返回') JSON"`
	CreateDate time.Time `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate time.Time `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}
