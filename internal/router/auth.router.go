package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b7-gin/internal/controller"
	"github.com/kodacampmain/koda-b7-gin/internal/repository"
	"github.com/kodacampmain/koda-b7-gin/internal/service"
)

func RegisterAuthRouter(router *gin.Engine, db *pgxpool.Pool) {
	authRouter := router.Group("/auth")

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	authRouter.POST("", authController.Login)
	authRouter.POST("/register", authController.Register)
}
