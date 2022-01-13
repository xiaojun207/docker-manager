package model

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"github.com/xiaojun207/go-base-utils/utils"
	"xorm.io/xorm"
)

type Page struct {
	boot.BindQuery
	CurrentPage int   `json:"currentPage" form:"currentPage"` // 绑定 query 需要"form" tag
	PageSize    int   `json:"pageSize" form:"pageSize"`       // 绑定 query 需要"form" tag
	Total       int64 `json:"total" form:"total"`             // 绑定 query 需要"form" tag
}

func (e *Page) IsUse() bool {
	return e.PageSize > 0
}

func (e *Page) FindPage(session *xorm.Session, record interface{}) (err error) {
	e.SetPageSql(session)
	e.Total, err = session.FindAndCount(record)
	return
}

func (e *Page) SetPageSql(session *xorm.Session) {
	if e.IsUse() {
		session.Limit(e.PageSize, (e.CurrentPage-1)*e.PageSize)
	}
}

func GetPage(c *gin.Context) (page Page) {
	currentPage := c.Query("currentPage")
	pageSize := c.Query("pageSize")
	if currentPage != "" && pageSize != "" {
		return Page{
			CurrentPage: utils.StrToInt(currentPage),
			PageSize:    utils.StrToInt(pageSize),
		}
	}

	return
}
