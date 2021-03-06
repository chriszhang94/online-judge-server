package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-judge-server/middleware"
	"online-judge-server/routers"
)

func InitRouter()*gin.Engine{
	Router := gin.Default()
	Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"success": true,
		})
	})

	Router.Use(middleware.Cors())
	ApiGroup := Router.Group("/api/v1")
	routers.InitProblemsRouter(ApiGroup)
	return Router
}
