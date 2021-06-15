package mgr

import (
	"docker-manager/data"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"docker-manager/web/ws"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"log"
	"strings"
)

func GetServers(c *gin.Context) {
	res := []interface{}{}

	data.Servers.ForEachMap(func(_ string, value map[string]interface{}) {
		server := value
		serverName := server["Name"].(string)
		if ws.IsConnected(serverName) {
			server["State"] = "connected"
		} else {
			server["State"] = "disconnect"
		}
		res = append(res, value)
	})
	resp.Resp(c, "100200", "成功", res)
}

func GetServerNames(c *gin.Context) {
	res := data.Servers.Keys()
	resp.Resp(c, "100200", "成功", res)
}

func GetContainers(c *gin.Context) {
	serverNames := c.QueryArray("serverNames[]")
	ContainerNames := c.QueryArray("ContainerNames[]")
	state := c.Query("state")
	log.Println("serverNames:", serverNames, "ContainerNames:", ContainerNames, ",state:", state)
	res := []interface{}{}

	data.Containers.ForEachArrMap(func(_ string, arr []map[string]interface{}) {
		for _, v := range arr {
			container := v

			c_state := container["State"].(string)
			c_serverName := container["ServerName"].(string)
			c_name := container["Name"].(string)

			if state != "" && c_state != state {
				continue
			}

			if len(serverNames) > 0 && !utils.StrInArr(serverNames, c_serverName) {
				continue
			}
			if len(ContainerNames) > 0 && !utils.StrInArr(ContainerNames, c_name) {
				continue
			}
			res = append(res, v)
		}
	})

	resp.Resp(c, "100200", "成功", res)
}

func GetContainerInfos(c *gin.Context) {
	serverNames := strings.Join(c.QueryArray("serverNames[]"), ",")
	state := c.Query("state")
	log.Println("serverName:", serverNames, ",state:", state)
	res := []interface{}{}

	data.Containers.ForEachArrMap(func(key string, arr []map[string]interface{}) {
		serverName := key

		containers := []map[string]string{}
		for _, v := range arr {
			container := v

			c_state := container["State"].(string)
			c_serverName := container["ServerName"].(string)

			if state != "" && c_state != state {
				continue
			}

			if len(serverNames) > 0 && !strings.Contains(serverNames, c_serverName) {
				continue
			}

			c := map[string]string{}
			c["ServerName"] = container["ServerName"].(string)
			c["Id"] = container["Id"].(string)

			c["Names"] = container["Name"].(string)
			c["State"] = container["State"].(string)
			containers = append(containers, c)
		}

		server := map[string]interface{}{}
		server["serverName"] = serverName
		server["containers"] = containers
		res = append(res, server)
	})

	resp.Resp(c, "100200", "成功", res)
}

func ContainerCmd(c *gin.Context) {
	//args := c.Query("args")
	//c.QueryArray("")
	args := []string{"docker", "start", "drone"}
	param := map[string]interface{}{
		"id":   uuid.New(),
		"args": args,
	}
	serverName := "docker-desktop"
	err := ws.Push(serverName, "base.cmd", param)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", "成功", err)
		return
	}
	resp.Resp(c, "100200", "成功", "")
}
