package web

import (
	"docker-manager/web/handler/agent"
	"docker-manager/web/handler/mgr"
	"docker-manager/web/handler/user"
	"docker-manager/web/ws"
	"github.com/gin-gonic/gin"
)

/**
路由控制
*/
var BaseRouter = func(router *gin.RouterGroup) {
	//router.Use(AuthInterceptor)
	//router.POST("/login", loginHandler)
	//router.GET("/hello", helloHandler)
	router.POST("/reg", AgentTokenInterceptor, agent.RegDockerHandler)
	router.POST("/containers", AgentTokenInterceptor, agent.ContainersHandler)
	router.POST("/containers/stats", AgentTokenInterceptor, agent.ContainersStatsHandler)
	router.POST("/container/stats", AgentTokenInterceptor, agent.ContainerStatsHandler)
	router.POST("/images", AgentTokenInterceptor, agent.ImagesHandler)
	router.POST("/image/", AgentTokenInterceptor, agent.ImagesHandler)
	router.GET("/config", AgentTokenInterceptor, agent.GetConfig)

	router.GET("/mgr/servers", AuthInterceptor, mgr.GetServers)
	router.GET("/mgr/serverNames", AuthInterceptor, mgr.GetServerNames)
	router.GET("/mgr/containers", AuthInterceptor, mgr.GetContainers)
	router.POST("/mgr/containers/update", AuthInterceptor, mgr.UpdateContainerList)
	router.GET("/mgr/stats", AuthInterceptor, mgr.GetStats)
	router.POST("/mgr/stats/update", AuthInterceptor, mgr.UpdateStats)
	router.GET("/mgr/containerInfos", AuthInterceptor, mgr.GetContainerInfos)
	router.POST("/mgr/container/:operator", AuthInterceptor, mgr.ContainerOperatorHandler)

	router.GET("/mgr/app/list", AuthInterceptor, mgr.ServiceList)
	router.POST("/mgr/app/del", AuthInterceptor, mgr.DeleteService)
	router.POST("/mgr/app/update", AuthInterceptor, mgr.UpdateService)
	router.GET("/mgr/app/group", AuthInterceptor, mgr.AppGroupList)

	router.GET("/mgr/dashboardSize", AuthInterceptor, mgr.GetDashboardSize)
	router.GET("/mgr/serverSize", AuthInterceptor, mgr.GetServerSize)
	router.GET("/mgr/conatainerSize", AuthInterceptor, mgr.GetContainerSize)
	router.GET("/mgr/imageSize", AuthInterceptor, mgr.GetImageSize)
	router.GET("/mgr/taskSize", AuthInterceptor, mgr.GetTaskSize)
	router.GET("/mgr/appSize", AuthInterceptor, mgr.GetAppSize)

	router.POST("/mgr/log/follows", AuthInterceptor, mgr.GetFollowLogList)
	router.POST("/mgr/log/start", AuthInterceptor, mgr.LogFollowStart)
	router.POST("/mgr/log/close", AuthInterceptor, mgr.LogFollowClose)

	router.POST("/mgr/publish", AuthInterceptor, mgr.PublishHandler)
	router.GET("/mgr/tasks", AuthInterceptor, mgr.GetTasks)
	router.GET("/mgr/cmd", AuthInterceptor, mgr.ContainerCmd)

	router.GET("/mgr/config", AuthInterceptor, mgr.GetConfig)
	router.POST("/mgr/config/update", AuthInterceptor, mgr.UpdateConfig)

	router.POST("/user/login", user.LoginHandler)
	router.POST("/user/logout", user.LogoutHandler)
	router.GET("/user/info", AuthInterceptor, user.UserInfoHandler)
	router.GET("/base/version", user.VersionHandler)

}

// 路由，把URL和执行方法连接起来
var WsRouter = func(router *gin.RouterGroup) {
	router.GET("/ws", AgentTokenInterceptor, ws.WSAgentHandler)
	router.GET("/ws/log", AuthInterceptor, ws.WSManagerHandler)
	//router.GET("/ws/:token", TokenInterceptor, wsHandler)
}
