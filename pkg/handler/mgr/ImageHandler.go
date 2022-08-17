package mgr

import (
	"docker-manager/pkg/data"
	"docker-manager/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func ImageList(c *gin.Context, req struct {
	ServerNames []string `json:"serverNames" form:"serverNames[]"`
	TagName     string   `json:"tagName" form:"tagName"`
}, page model.Page) {
	log.Println("req.ServerNames:", req.ServerNames)
	imageList, total, err := data.GetImages(req.ServerNames, req.TagName, page)
	page.Total = total
	if err != nil {
		log.Println(err)
		boot.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	var list []map[string]interface{}
	for _, image := range imageList {
		list = append(list, map[string]interface{}{
			"ServerName":  image.ServerName,
			"ImageId":     image.ImageId,
			"RepoTags":    image.RepoTags,
			"Size":        image.Size,
			"RepoDigests": image.RepoDigests,
			"Created":     image.Created,
		})
	}
	boot.Resp(c, "100200", "成功", gin.H{
		"list": list,
		"page": page,
	})
}

func ImageDetail(c *gin.Context, req struct {
	ImageId string `json:"ImageId" form:"ImageId"`
}) {
	if req.ImageId == "" {
		boot.Resp(c, "100100", "字段校验错误", "")
		return
	}
	res, err := data.GetImage(req.ImageId)
	if err != nil {
		boot.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	boot.Resp(c, "100200", "成功", res)
}
