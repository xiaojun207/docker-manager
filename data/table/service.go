package table

import (
	"time"
)

type Service struct {
	Id         int                      `xorm:"not null pk autoincr INT"`
	Name       string                   `xorm:"comment('服务名称（同容器名称）') VARCHAR(40) unique"`
	Image      string                   `xorm:"comment('镜像名称') VARCHAR(1024)"`
	Ports      []map[string]interface{} `xorm:"comment('端口列表') JSON"`
	Env        []map[string]interface{} `xorm:"comment('环境变量列表') JSON"`
	Vol        []map[string]interface{} `xorm:"comment('路径映射列表') JSON"`
	Memory     int64                    `xorm:"comment('内存限制') BIGINT"`
	Running    int                      `xorm:"default 0 comment('运行实例数量') INT"`
	Replicas   int                      `xorm:"default 0 comment('副本数量') INT"`
	Status     int                      `xorm:"default 0 comment('状态，0 不在线，1在线') INT"`
	Summary    map[string]interface{}   `xorm:"JSON"`
	CreateDate time.Time                `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate time.Time                `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}
