package table

import "time"

type Image struct {
	Id         int                    `xorm:"not null pk autoincr INT"`
	ServerName string                 `xorm:"comment('所在服务器名称') VARCHAR(40)"`
	Created    int64                  `xorm:"default 0 comment('Created时间') INT"`
	Image      string                 `xorm:"comment('镜像') VARCHAR(240)"`
	ImageID    string                 `xorm:"comment('镜像Id') VARCHAR(140)"`
	Status     string                 `xorm:"default '' comment('状态2') VARCHAR(24)"`
	Size       int64                  `xorm:"default 0 comment('空间占用') BIGINT"`
	Summary    map[string]interface{} `xorm:"JSON"`
	CreateDate time.Time              `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate time.Time              `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}
