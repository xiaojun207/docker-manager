package mgr

import (
	"docker-manager/data"
	"docker-manager/model"
	"github.com/gin-gonic/gin"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func ImageList(c *gin.Context, page model.Page) {
	serverNames := c.QueryArray("serverNames[]")
	tagName := c.Query("tagName")
	//log.Println("page:", page)
	imageList, total, err := data.GetImages(serverNames, tagName, page)
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

func ImageDetail(c *gin.Context) {
	imageId := c.Query("ImageId")
	if imageId == "" {
		boot.Resp(c, "100100", "字段校验错误", "")
		return
	}
	res, err := data.GetImage(imageId)
	if err != nil {
		boot.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	boot.Resp(c, "100200", "成功", res)
}
