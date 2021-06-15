package mgr

import (
	"docker-manager/data"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetStats(c *gin.Context) {
	serverNames := c.QueryArray("serverNames[]")
	ContainerNames := c.QueryArray("ContainerNames[]")
	Follow := c.Query("Follow")
	log.Println("serverNames:", serverNames, ",ContainerNames:", ContainerNames, ",Follow:", Follow)
	res := []interface{}{}

	data.Stats.ForEachArr(func(ServerName string, arr []interface{}) {
		if len(serverNames) > 0 && !utils.StrInArr(serverNames, ServerName) {
			return
		}

		for _, v := range arr {
			container := v.(map[string]interface{})
			c_name := utils.TrimContainerName(container["name"])
			container["Name"] = c_name

			if len(ContainerNames) > 0 && !utils.StrInArr(ContainerNames, c_name) {
				continue
			}
			if len(Follow) > 0 {
				f, _ := strconv.ParseBool(Follow)
				if f != container["Follow"].(bool) {
					continue
				}
			}
			res = append(res, container)
		}
	})

	resp.Resp(c, "100200", "成功", res)
}
