package utils

import "github.com/gin-gonic/gin"

func GetRemoteIP(c *gin.Context) string {
	ip := c.GetHeader("X-Real-IP")
	if ip == "" {
		ip = c.ClientIP()
	}
	return ip
}
