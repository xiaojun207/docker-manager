package mgr

import (
	"docker-manager/data"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
)

func ImageList(c *gin.Context) {
	serverNames := c.QueryArray("serverNames[]")
	tagName := c.Query("tagName")
	imageList, err := data.GetImages(serverNames, tagName)
	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	var res []map[string]interface{}
	for _, image := range imageList {

		if len(serverNames) > 0 && !utils.StrInArr(serverNames, image.ServerName) {
			continue
		}

		res = append(res, map[string]interface{}{
			"ServerName":  image.ServerName,
			"ImageId":     image.ImageId,
			"RepoTags":    image.RepoTags,
			"Size":        image.Size,
			"RepoDigests": image.RepoDigests,
			"Created":     image.Created,
		})
	}
	resp.Resp(c, "100200", "成功", res)
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
