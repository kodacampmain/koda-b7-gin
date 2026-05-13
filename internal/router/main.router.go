package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b7-gin/internal/dto"
)

func InitRouter(router *gin.Engine) {
	// implementasi route
	// router.METHOD(endpoint, callback)
	RegisterRootRouter(router)
	RegisterUserRouter(router)
	RegisterMovieRouter(router)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Message: "Invalid Route",
			Success: false,
			Error:   "route not found",
		})
	})
}
