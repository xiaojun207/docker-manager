package web

import (
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RespApi(c *gin.Context, res resp.ApiResp) {
	c.JSON(http.StatusOK, gin.H{
		"code": res.Code,
		"msg":  res.Msg,
		"data": res.Data,
	})
}

func RespForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"code": "403", "msg": "禁止访问",
	})
}

func RespSuccessNoData(c *gin.Context) {
	resp.Resp(c, "100200", "成功", "")
}

func RespSuccess(c *gin.Context, data interface{}) {
	resp.Resp(c, "100200", "成功", data)
}

func _404Handler(c *gin.Context) {
	resp.Resp(c, "100104", "无效的请求", "")
}

func Start(port string, contextPath string, routers ...func(router *gin.RouterGroup)) {
	gin.SetMode(gin.ReleaseMode)
	rootRouter := gin.New()
	rootRouter.Use(gin.Logger())
	rootRouter.Use(gin.Recovery())
	rootRouter.NoRoute(_404Handler)
	rootRouter.StaticFile("/", "./views/index.html")
	rootRouter.StaticFile("/favicon.ico", "./views/favicon.ico")
	rootRouter.StaticFile("/index-batch.html", "./views/index-batch.html")
	rootRouter.StaticFS("/static/", http.Dir("./views/static/"))
	//rootRouter.Static("/", "./views/")
	apiRouter := rootRouter.Group(contextPath)
	BaseRouter(apiRouter)
	WsRouter(apiRouter)

	if routers != nil {
		for _, router := range routers {
			router(apiRouter)
		}
	}

	//l, err := net.Listen("tcp", "127.0.0.1:" + port)
	//if err != nil {
	//	fmt.Println("Listen: %v", err)
	//}
	//l = netutil.LimitListener(l, 1)
	//http.Serve(l, rootRouter)

	log.Println("Web Server starting http://127.0.0.1:" + port + contextPath)
	http.ListenAndServe(":"+port, rootRouter)

}
