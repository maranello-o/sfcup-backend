package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sfcup/handler"
	"sfcup/middleware"
)

func EngineStart() {
	engine := gin.Default()
	//测试请求
	engine.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	user := engine.Group("/user")
	{
		user.GET("/code", handler.GenVerificationCode)
		user.POST("registry", handler.RegisterUser)
		needAuthUser := user.Use(middleware.CheckAuth)
		needAuthUser.POST("/login", handler.Login)
	}
	model := engine.Group("/model")
	{
		model.POST("/prediction", handler.GetPredictResult)
	}

	if err := engine.Run("0.0.0.0:8088"); err != nil {
		fmt.Println(err)
		return
	}
}
