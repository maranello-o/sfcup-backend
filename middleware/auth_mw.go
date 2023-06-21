package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sfcup/response"
	"sfcup/util"
)

func CheckAuth(c *gin.Context) {
	token := c.GetHeader("sfcup_token")
	//fmt.Println(token, c.Request.Header)
	if token == "" {
		response.Send(c, http.StatusUnauthorized, nil, "无权限")
		c.Abort()
		return
	}
	claims, err := util.ParseJWT(token)
	if err != nil {
		response.Send(c, http.StatusUnauthorized, nil, "无权限")
		c.Abort()
		return
	}
	c.Set("id", claims.ID)

	c.Next()
}
