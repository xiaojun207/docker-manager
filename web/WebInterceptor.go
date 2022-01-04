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
	reqIP := utils.GetRemoteIP(c)
	if !service.IsWhiteIp(reqIP) {
		log.Println("禁止访问,IP:", reqIP)
		resp.Resp(c, "403", "禁止访问", "")
		c.Status(403)
		c.Abort()
		return
	}
}

func AgentTokenInterceptor(c *gin.Context) {
	token := c.GetHeader("authorization")
	if token == "" {
		reqIP := utils.GetRemoteIP(c)
		log.Println("URI:", c.Request.RequestURI+", AgentTokenInterceptor token is empty:", ",URI:", c.Request.RequestURI, ",fromIp:", reqIP)
		resp.Resp(c, "105101", "账户未登录", "")
		c.Abort()
		return
	}
	isValid, uid, err := utils.ValidToken(token)
	if err != nil || !isValid {
		reqIP := utils.GetRemoteIP(c)
		log.Println("AgentToken.err:", err, ", isValid:", isValid, ",URI:", c.Request.RequestURI, ",fromIp:", reqIP)
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
		reqIP := utils.GetRemoteIP(c)
		log.Println("ApiTokenInterceptor:", apiKey, apiSecret, ",URI:", c.Request.RequestURI, ",fromIp:", reqIP)
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
			reqIP := utils.GetRemoteIP(c)
			log.Println("AuthInterceptor.err", err, ",URI:", c.Request.RequestURI, ",fromIp:", reqIP)
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
		reqIP := utils.GetRemoteIP(c)
		log.Println("AuthInterceptor.err:", err, ", isValid:", isValid, ",URI:", c.Request.RequestURI, ",fromIp:", reqIP)
		resp.Resp(c, "105101", "账户未登录", "")
		c.Abort()
		return
	}

	id, err := strconv.Atoi(uid)
	c.Set("uid", id)
}
