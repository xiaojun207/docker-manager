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

	//
	router.POST("/agent/reg", AgentTokenInterceptor, agent.RegDockerHandler)
	router.POST("/agent/containers", AgentTokenInterceptor, agent.ContainersHandler)
	router.POST("/agent/containers/stats", AgentTokenInterceptor, agent.ContainersStatsHandler)
	router.POST("/agent/container/stats", AgentTokenInterceptor, agent.ContainerStatsHandler)
	router.POST("/agent/images", AgentTokenInterceptor, agent.ImagesHandler)
	router.POST("/agent/image/", AgentTokenInterceptor, agent.ImagesHandler)
	router.GET("/agent/config", AgentTokenInterceptor, agent.GetConfig)
	router.POST("/agent/login", agent.LoginHandler)

	///

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

	router.GET("/base/version", user.VersionHandler)

	router.POST("/user/login", user.LoginHandler)
	router.POST("/user/logout", user.LogoutHandler)
	router.GET("/user/info", AuthInterceptor, user.UserInfoHandler)
	router.GET("/user/userList", AuthInterceptor, user.UserListHandler)
	router.POST("/user/alterPassword", AuthInterceptor, user.AlterPasswordHandler)
	router.POST("/user/resetPassword", AuthInterceptor, user.ResetPasswordHandler)
	router.POST("/user/changeStatus", AuthInterceptor, user.ChangeStatusHandler)
	router.POST("/user/addUser", AuthInterceptor, user.AddUserHandler)
	router.POST("/user/deleteUser", AuthInterceptor, user.DeleteUserHandler)

}

// 路由，把URL和执行方法连接起来
var WsRouter = func(router *gin.RouterGroup) {
	router.GET("/agent/ws", AgentTokenInterceptor, ws.WSAgentHandler)
	router.GET("/ws/log", AuthInterceptor, ws.WSManagerHandler)
	//router.GET("/ws/:token", TokenInterceptor, wsHandler)
}
