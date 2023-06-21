package response

import "github.com/gin-gonic/gin"

func Send(context *gin.Context, code int, data any, msg string) {
	context.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
