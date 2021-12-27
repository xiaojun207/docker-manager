package web

import (
	"docker-manager/service"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func WhiteIpInterceptor(c *gin.Context) {
	reqIP := c.GetHeader("HostIp")
	if reqIP == "" {
		reqIP = c.ClientIP()
	}

	if !service.IsWhiteIp(reqIP) {
		resp.Resp(c, "403", "禁止访问", "")
		c.Status(403)
		c.Abort()
		return
	}
}

func AgentTokenInterceptor(c *gin.Context) {
	token := c.GetHeader("authorization")
	if token == "" {
		log.Println("URI:", c.Request.RequestURI+", AgentTokenInterceptor token is empty:")
		resp.Resp(c, "105101", "账户未登录", "")
		c.Abort()
		return
	}
	isValid, uid, err := utils.ValidToken(token)
	if err != nil || !isValid {
		log.Println("AgentToken.err:", err, ", isValid:", isValid)
		resp.Resp(c, "105101", "账户未登录", "")
		c.Abort()
		return
	}

	id, err := strconv.Atoi(uid)
	c.Set("uid", id)
}

func ApiTokenInterceptor(c *gin.Context) {
	apiKey := c.GetHeader("authorization")
	apiSecret := c.GetHeader("authorization")
	if apiKey == "" || !service.LoginApi(apiKey, apiSecret) {
		log.Println("URI:", c.Request.RequestURI+", ApiTokenInterceptor:", apiKey, apiSecret)
		resp.Resp(c, "105101", "账户未登录", "")
		c.Abort()
		return
	}
}

func AuthInterceptor(c *gin.Context) {
	token := c.GetHeader("authorization")
	if token == "" {
		_token, err := c.Request.Cookie("c-token")
		if err != nil {
			log.Println("AuthInterceptor.err", err)
			resp.Resp(c, "105101", "账户未登录", "")
			c.Abort()
			return
		}
		token = _token.Value
	}

	if token == "" {
		resp.Resp(c, "105101", "账户未登录", "")
		c.Abort()
		return
	}

	isValid, uid, err := utils.ValidToken(token)
	if err != nil || !isValid {
		log.Println("AuthInterceptor.err:", err, ", isValid:", isValid)
		resp.Resp(c, "105101", "账户未登录", "")
		c.Abort()
		return
	}

	id, err := strconv.Atoi(uid)
	c.Set("uid", id)
}
