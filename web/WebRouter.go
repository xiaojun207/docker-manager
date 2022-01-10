package web

import (
	"docker-manager/web/handler/agent"
	"docker-manager/web/handler/base"
	"docker-manager/web/handler/mgr"
	"docker-manager/web/handler/user"
	"docker-manager/web/ws"
	"github.com/xiaojun207/gin-boot/boot"
	"net/http"
)

var ApiRouter = func(router *boot.WebRouter) {
	router.Use(WhiteIpInterceptor)

	// base
	baseRouter(router)
	wsRouter(router)
	agentRouter(router.Group("/agent"))
	mgrRouter(router.Group("/mgr"))
	userRouter(router.Group("/user"))
}

//////////////////////////////////////////////////////////////////////////////////////////////////

var StaticRouter = func(router *boot.WebRouter) {
	router.StaticFile("/", "./views/index.html")
	router.StaticFile("/favicon.ico", "./views/favicon.ico")
	router.StaticFS("/static/", http.Dir("./views/static/"))
}

/////////////

var baseRouter = func(router *boot.WebRouter) {
	router.GET("/base/version", base.VersionHandler)
	router.GET("/base/Textversion", base.VersionTextHandler)
	router.POST("/base/sendCode", base.SendCodeHandler)
	router.GET("/base/dbtrace", AuthInterceptor, base.DBTraceHandler)
	router.GET("/base/forbidden/log", AuthInterceptor, base.ForbiddenLogHandler)
	router.GET("/test", base.TestHandler)
}

// WsRouter 路由，把URL和执行方法连接起来
var wsRouter = func(router *boot.WebRouter) {
	router.GET("/agent/ws", AgentTokenInterceptor, ws.WSAgentHandler)
	router.GET("/ws/log", AuthInterceptor, ws.WSManagerHandler)
	router.GET("/ws/exec", AuthInterceptor, ws.WSManagerExecHandler)
}

var userRouter = func(router *boot.WebRouter) {
	router.POST("/login", user.LoginHandler)
	router.POST("/logout", user.LogoutHandler)
	router.POST("/forgetPassword", user.ForgetPasswordHandler)

	router.GET("/info", AuthInterceptor, user.UserInfoHandler)
	router.GET("/userList", AuthInterceptor, user.UserListHandler)
	router.POST("/alterPassword", AuthInterceptor, user.AlterPasswordHandler)
	router.POST("/resetPassword", AuthInterceptor, user.ResetPasswordHandler)
	router.POST("/changeStatus", AuthInterceptor, user.ChangeStatusHandler)
	router.POST("/addUser", AuthInterceptor, user.AddUserHandler)
	router.POST("/deleteUser", AuthInterceptor, user.DeleteUserHandler)
}

var agentRouter = func(router *boot.WebRouter) {
	//router.Use(AuthInterceptor)
	router.POST("/reg", AgentTokenInterceptor, agent.RegDockerHandler)
	router.POST("/containers", AgentTokenInterceptor, agent.ContainersHandler)
	router.POST("/containers/stats", AgentTokenInterceptor, agent.ContainersStatsHandler)
	router.POST("/container/stats", AgentTokenInterceptor, agent.ContainerStatsHandler)
	router.POST("/images", AgentTokenInterceptor, agent.ImagesHandler)
	router.POST("/image/", AgentTokenInterceptor, agent.ImagesHandler)
	router.GET("/config", AgentTokenInterceptor, agent.GetConfig)
	router.POST("/login", agent.LoginHandler)
}

var mgrRouter = func(router *boot.WebRouter) {

	router.GET("/server/list", AuthInterceptor, mgr.GetServers)
	router.GET("/server/detail", AuthInterceptor, mgr.GetServer)
	router.GET("/server/names", AuthInterceptor, mgr.GetServerNames)
	router.POST("/server/delete", AuthInterceptor, mgr.DeleteServer)

	router.GET("/container/list", AuthInterceptor, mgr.GetContainers)
	router.POST("/containers/update", AuthInterceptor, mgr.UpdateContainerList)
	router.GET("/containerInfos", AuthInterceptor, mgr.GetContainerInfos)
	router.GET("/container/detail", AuthInterceptor, mgr.GetContainer)
	router.POST("/container/:operator", AuthInterceptor, mgr.ContainerOperatorHandler)

	router.GET("/stats/list", AuthInterceptor, mgr.GetStatsList)
	router.GET("/stats/detail", AuthInterceptor, mgr.GetStats)
	router.POST("/stats/update", AuthInterceptor, mgr.UpdateStats)

	router.GET("/image/list", AuthInterceptor, mgr.ImageList)
	router.GET("/image/detail", AuthInterceptor, mgr.ImageDetail)
	router.POST("/image/:operator", AuthInterceptor, mgr.ImageCmd)

	router.GET("/app/list", AuthInterceptor, mgr.ServiceList)
	router.POST("/app/del", AuthInterceptor, mgr.DeleteService)
	router.POST("/app/update", AuthInterceptor, mgr.UpdateService)
	router.GET("/app/group", AuthInterceptor, mgr.AppGroupList)
	router.POST("/app/group/del", AuthInterceptor, mgr.DeleteGroup)

	router.GET("/dashboardSize", AuthInterceptor, mgr.GetDashboardSize)
	router.GET("/serverSize", AuthInterceptor, mgr.GetServerSize)
	router.GET("/conatainerSize", AuthInterceptor, mgr.GetContainerSize)
	router.GET("/imageSize", AuthInterceptor, mgr.GetImageSize)
	router.GET("/taskSize", AuthInterceptor, mgr.GetTaskSize)
	router.GET("/appSize", AuthInterceptor, mgr.GetAppSize)

	router.POST("/log/follows", AuthInterceptor, mgr.GetFollowLogList)
	router.POST("/log/start", AuthInterceptor, mgr.LogFollowStart)
	router.POST("/log/close", AuthInterceptor, mgr.LogFollowClose)

	router.POST("/publish", AuthInterceptor, mgr.PublishHandler)
	router.POST("/publish/yaml", AuthInterceptor, mgr.PublishYamlHandler)
	router.GET("/tasks", AuthInterceptor, mgr.GetTasks)
	router.POST("/task/del", AuthInterceptor, mgr.DelTask)
	router.GET("/cmd", AuthInterceptor, mgr.ContainerCmd)

	router.GET("/config", AuthInterceptor, mgr.GetConfig)
	router.POST("/config/update", AuthInterceptor, mgr.UpdateConfig)
	router.GET("/config/whiteList", AuthInterceptor, mgr.GetWhiteList)
	router.POST("/config/addWhiteIp", AuthInterceptor, mgr.AddWhiteIp)
	router.POST("/config/deleteWhiteIp", AuthInterceptor, mgr.DeleteWhiteIp)
}
