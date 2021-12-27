package mgr

import (
	"docker-manager/data"
	"docker-manager/model"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
)

func ImageList(c *gin.Context) {
	serverNames := c.QueryArray("serverNames[]")
	tagName := c.Query("tagName")
	page := model.GetPage(c)
	//log.Println("page:", page)
	imageList, total, err := data.GetImages(serverNames, tagName, page)
	page.Total = total
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	var list []map[string]interface{}
	for _, image := range imageList {

		if len(serverNames) > 0 && !utils.StrInArr(serverNames, image.ServerName) {
			continue
		}

		list = append(list, map[string]interface{}{
			"ServerName":  image.ServerName,
			"ImageId":     image.ImageId,
			"RepoTags":    image.RepoTags,
			"Size":        image.Size,
			"RepoDigests": image.RepoDigests,
			"Created":     image.Created,
		})
	}
	resp.Resp(c, "100200", "成功", gin.H{
		"list": list,
		"page": page,
	})
}

func ImageDetail(c *gin.Context) {
	imageId := c.Query("ImageId")
	if imageId == "" {
		resp.Resp(c, "100100", "字段校验错误", "")
		return
	}
	res, err := data.GetImage(imageId)
	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	resp.Resp(c, "100200", "成功", res)
}
