package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b7-gin/internal/controller"
)

func RegisterRootRouter(router *gin.Engine) {
	rootRouter := router.Group("/")

	rootController := controller.NewRootController()

	rootRouter.GET("", rootController.HelloKoda)
	rootRouter.POST("", rootController.HelloString)
}
