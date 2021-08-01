package mgr

import (
	"docker-manager/data"
	"docker-manager/service"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"docker-manager/web/ws"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"log"
	"strings"
)

func GetServers(c *gin.Context) {
	servers, _ := data.GetServers()
	for _, server := range servers {
		if ws.AgentConnected(server.Name) {
			server.State = "connected"
		} else {
			server.State = "disconnect"
		}
	}
	resp.Resp(c, "100200", "成功", servers)
}

func GetServerNames(c *gin.Context) {
	res, _ := data.GetServersName()
	resp.Resp(c, "100200", "成功", res)
}

func UpdateContainerList(c *gin.Context) {
	ch := "docker.container.list"
	service.SendToAllServer(ch, map[string]interface{}{})
	resp.Resp(c, "100200", "成功", "")
}

func GetContainers(c *gin.Context) {
	serverNames := strings.Join(c.QueryArray("serverNames[]"), ",")
	ContainerNames := c.QueryArray("ContainerNames[]")
	state := c.Query("state")
	log.Println("serverNames:", serverNames, "ContainerNames:", ContainerNames, ",state:", state)
	res := []interface{}{}

	containers, err := data.GetContainers()
	log.Println("GetContainers.err:", err)

	for _, container := range containers {
		if state != "" && container.State != state {
			continue
		}
		if len(serverNames) > 0 && !strings.Contains(serverNames, container.ServerName) {
			continue
		}
		if len(ContainerNames) > 0 && !utils.StrInArr(ContainerNames, container.Name) {
			continue
		}
		res = append(res, container)
	}
	resp.Resp(c, "100200", "成功", res)
}

// info,服务和容器基本信息
func GetContainerInfos(c *gin.Context) {
	serverNames := strings.Join(c.QueryArray("serverNames[]"), ",")
	state := c.Query("state")
	log.Println("serverName:", serverNames, ",state:", state)

	res := []interface{}{}
	containers, err := data.GetContainers()
	log.Println("GetContainerInfos.err:", err)

	serverMap := map[string]map[string]interface{}{}
	for _, container := range containers {
		if state != "" && container.State != state {
			continue
		}
		if len(serverNames) > 0 && !strings.Contains(serverNames, container.ServerName) {
			continue
		}

		server, ok := serverMap[container.ServerName]
		var containers []map[string]string
		if ok {
			containers = server["containers"].([]map[string]string)
		} else {
			server = map[string]interface{}{}
			containers = []map[string]string{}
			server["serverName"] = container.ServerName
			server["containers"] = containers
			serverMap[container.ServerName] = server
			res = append(res, server)
		}

		c := map[string]string{}
		c["ServerName"] = container.ServerName
		c["Id"] = container.ContainerId
		c["Names"] = container.Name
		c["State"] = container.State
		containers = append(containers, c)
	}

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
	err := ws.AgentWsConnectGroup.Push(serverName, "base.cmd", param)
	if err != nil {
		log.Println(err)
		resp.Resp(c, "100100", "成功", err)
		return
	}
	resp.Resp(c, "100200", "成功", "")
}
