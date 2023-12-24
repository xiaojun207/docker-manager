package mgr

import (
	"docker-manager/pkg/data"
	"docker-manager/pkg/model"
	"docker-manager/pkg/service"
	"docker-manager/pkg/ws"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/xiaojun207/gin-boot/boot"
	"log"
)

func UpdateContainerList(c *gin.Context) {
	ch := "docker.container.list"
	service.SendToAllServer(ch, map[string]interface{}{})
	boot.Resp(c, boot.CodeSuccess, "成功", "")
}

func GetContainers(c *gin.Context, page model.Page, req struct {
	ServerNames    []string `json:"serverNames" form:"ServerNames[]"`
	ContainerNames []string `json:"serverNames" form:"ContainerNames[]"`
	State          string   `json:"state" form:"state"`
}) {
	log.Println("GetContainers.req:", req)
	containers, err := data.GetContainers(req.State, req.ServerNames, req.ContainerNames, &page)
	if err != nil {
		log.Println("GetContainers.err:", err)
	}

	var res []map[string]interface{}
	for _, container := range containers {
		res = append(res, map[string]interface{}{
			"ServerName":  container.ServerName,
			"ContainerId": container.ContainerId,
			"Name":        container.Name,
			"Image":       container.Image,
			"Status":      container.Status,
			"State":       container.State,
			"Ports":       container.Ports,
			"Created":     container.Created,
		})
	}
	boot.Resp(c, boot.CodeSuccess, "成功", gin.H{
		"list": res,
		"page": page,
	})
}

func GetContainer(c *gin.Context) {
	ContainerId := c.Query("ContainerId")
	container, err := data.GetContainer(ContainerId)
	log.Println("GetContainers.err:", err)
	boot.Resp(c, boot.CodeSuccess, "成功", container)
}

// info,服务和容器基本信息
func GetContainerInfos(c *gin.Context, req struct {
	ServerNames []string `json:"serverNames" form:"serverNames[]"`
	State       string   `json:"state" form:"state"`
}) {
	log.Println("GetContainerInfos.req:", req)
	containers, err := data.GetContainers(req.State, req.ServerNames, nil, &model.Page{})
	if err != nil {
		log.Println("GetContainerInfos.err:", err)
		return
	}

	serverMap := map[string]map[string]interface{}{}
	for _, container := range containers {
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

	var res []interface{}
	for _, server := range serverMap {
		res = append(res, server)
	}

	boot.Resp(c, boot.CodeSuccess, "成功", res)
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
		boot.Resp(c, "100100", "成功", err)
		return
	}
	boot.Resp(c, boot.CodeSuccess, "成功", "")
}
