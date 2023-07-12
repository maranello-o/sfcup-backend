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
	engine.Use(middleware.Cors)
	//测试请求
	engine.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	user := engine.Group("/user")
	{
		user.GET("/code", handler.GenVerificationCode)
		user.POST("/registry", handler.RegisterUser)
		user.POST("/login", handler.Login)
		needAuthUser := user.Use(middleware.CheckAuth)
		{
			needAuthUser.GET("/profile", handler.GetSelfProfile)
		}
	}
	model := engine.Group("/model")
	{
		needAuthModel := model.Use(middleware.CheckAuth)
		needAuthModel.POST("/:modelName/prediction/:fileName", handler.GetPredictResult)
	}
	image := engine.Group("/image")
	{
		image.GET("/count")
		image.POST("/volume", handler.GetVolume)
	}
	file := engine.Group("/file")
	{
		file.GET("/:fileName", handler.DownloadFile)
		needAuthFile := file.Use(middleware.CheckAuth)
		needAuthFile.POST("", handler.UploadFile)
		needAuthFile.GET("", handler.GetSelfFiles)
	}
	usage := engine.Group("/usage")
	{
		usage.GET("/proportion", handler.GetModelUsageProportion)
		needAuthUsage := usage.Use(middleware.CheckAuth)
		needAuthUsage.GET("", handler.GetSelfUsage)
		usage.GET("/statistic", handler.GetTotalUsageStatistic)
	}
	if err := engine.Run("0.0.0.0:8088"); err != nil {
		fmt.Println(err)
		return
	}
}
