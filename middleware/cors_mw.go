package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors(c *gin.Context) {
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin")
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", origin)
		//主要设置Access-Control-Allow-Origin
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Cookie,sfcup_token,Content-Type,Content-Length, Authorization, Accept,X-Requested-With,domain,zdy")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Set("content-type", "application/json")
	}
	//预检请求处理
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
