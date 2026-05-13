package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b7-gin/internal/dto"
	"github.com/kodacampmain/koda-b7-gin/internal/middleware"
)

func InitRouter(router *gin.Engine, db *pgxpool.Pool) {
	// middleware global
	// router.Use(middleware.Logger)
	router.Use(middleware.CORSMiddleware)
	// implementasi route
	// router.METHOD(endpoint, callback)
	RegisterRootRouter(router)
	RegisterUserRouter(router, db)
	RegisterMovieRouter(router)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Message: "Invalid Route",
			Success: false,
			Error:   "route not found",
		})
	})
}
