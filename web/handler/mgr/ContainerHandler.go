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
	servers, err := data.GetServers()

	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
	for i, server := range servers {
		if ws.AgentConnected(server.Name) {
			server.State = "connected"
		} else {
			server.State = "disconnect"
		}
		servers[i] = server
	}
	resp.Resp(c, "100200", "成功", servers)
}

func GetServerNames(c *gin.Context) {
	res, err := data.GetServersName()

	if err != nil {
		resp.Resp(c, "100100", "请求异常", err.Error())
		return
	}
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

	containers, err := data.GetContainers()
	if err != nil {
		log.Println("GetContainerInfos.err:", err)
	} else {
		log.Println("GetContainerInfos.len:", len(containers))
	}

	serverMap := map[string]map[string]interface{}{}
	for _, container := range containers {
		if state != "" && container.State != state {
			continue
		}
		if len(serverNames) > 0 && !strings.Contains(serverNames, container.ServerName) {
			continue
		}

		server, ok := serverMap[container.ServerName]
		var clist []map[string]string
		if ok {
			clist = server["containers"].([]map[string]string)
		} else {
			clist = []map[string]string{}
			server = map[string]interface{}{}
			server["serverName"] = container.ServerName
		}

		tmp := map[string]string{}
		tmp["ServerName"] = container.ServerName
		tmp["Id"] = container.ContainerId
		tmp["Name"] = container.Name
		tmp["State"] = container.State
		clist = append(clist, tmp)

		server["containers"] = clist
		serverMap[container.ServerName] = server
	}

	res := []interface{}{}
	for _, server := range serverMap {
		res = append(res, server)
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
