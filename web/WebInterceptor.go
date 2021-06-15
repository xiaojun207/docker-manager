package web

import (
	"docker-manager/data"
	"docker-manager/utils"
	"docker-manager/web/resp"
	"github.com/gin-gonic/gin"
	"log"
)

func AgentTokenInterceptor(c *gin.Context) {
	Token := data.SampleToken.GetStr("agentToken")
	token := c.Request.Header.Get("authorization")
	if token == "" || token != Token {
		log.Println("URI:", c.Request.RequestURI+", AgentTokenInterceptor:", token)
		resp.Resp(c, "105101", "账户未登录", "")
		c.Abort()
		return
	}
}

func ApiTokenInterceptor(c *gin.Context) {
	Token := data.SampleToken.GetStr("apiToken")
	token := c.Request.Header.Get("authorization")
	if token == "" || token != Token {
		log.Println("URI:", c.Request.RequestURI+", ApiTokenInterceptor:", token)
		resp.Resp(c, "105101", "账户未登录", "")
		c.Abort()
		return
	}
}

func AuthInterceptor(c *gin.Context) {
	token := c.Request.Header.Get("authorization")
	if token == "" {
		xxx_token, err := c.Request.Cookie("xxx_token")
		if err != nil {
			log.Println("AuthInterceptor.err", err)
			resp.Resp(c, "105101", "账户未登录", "")
			c.Abort()
			return
		}
		token = xxx_token.Value
	}

	if token == "" {
		resp.Resp(c, "105101", "账户未登录", "")
		c.Abort()
		return
	}

	isValid, data, err := utils.ValidToken(token)
	if err != nil {
		resp.Resp(c, "105101", "账户未登录1", "")
		c.Abort()
		return
	}
	if !isValid {
		resp.Resp(c, "105101", "账户未登录2", "")
		c.Abort()
		return
	}

	c.Set("username", data)
}
