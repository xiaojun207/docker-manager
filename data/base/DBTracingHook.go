package base

import (
	"context"
	"docker-manager/model"
	"xorm.io/xorm/contexts"
)

type TracingHook struct {
	BeforeNum int           `json:"before_num"`
	AfterNum  int           `json:"after_num"`
	SqlMap    model.SyncMap `json:"sql_map"`
}

var DBTracingHook = &TracingHook{}

// xorm的hook接口需要满足BeforeProcess和AfterProcess函数
func (h *TracingHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	h.BeforeNum++
	//log.Println("before.c.sql:", h.BeforeNum)
	h.SqlMap.IncInt(c.SQL, 1)
	return c.Ctx, nil
}

func (h *TracingHook) AfterProcess(c *contexts.ContextHook) error {
	h.AfterNum++
	//log.Println("after.c.sql:", h.afterNum)
	return nil
}
