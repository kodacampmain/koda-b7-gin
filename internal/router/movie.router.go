package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b7-gin/internal/controller"
)

func RegisterMovieRouter(router *gin.Engine) {
	movieRouter := router.Group("/movies")

	moviesController := controller.NewMoviesController()

	// query
	movieRouter.GET("", moviesController.Search)
}
