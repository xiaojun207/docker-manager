package table

import "time"

type Image struct {
	Id          int                    `xorm:"not null pk autoincr INT"`
	AppId       string                 `xorm:"comment('AppId') VARCHAR(40)" json:"AppId"`
	ServerName  string                 `xorm:"comment('所在服务器名称') VARCHAR(40)" json:"ServerName"`
	Created     int64                  `xorm:"default 0 comment('Created时间') INT" json:"Created"`
	RepoDigests string                 `xorm:"comment('镜像Digests') VARCHAR(240)" json:"RepoDigests"`
	RepoTags    string                 `xorm:"comment('镜像Tag') VARCHAR(540)" json:"RepoTags"`
	ImageId     string                 `xorm:"comment('镜像Id') VARCHAR(140)" json:"image_id"`
	Status      string                 `xorm:"default '' comment('状态2') VARCHAR(24)" json:"status"`
	Size        int64                  `xorm:"default 0 comment('空间占用') BIGINT" json:"Size"`
	Summary     map[string]interface{} `xorm:"JSON" json:"Summary"`
	CreateDate  time.Time              `xorm:"created default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateDate  time.Time              `xorm:"updated default CURRENT_TIMESTAMP TIMESTAMP"`
}
