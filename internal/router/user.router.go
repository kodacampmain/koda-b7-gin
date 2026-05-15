package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b7-gin/internal/controller"
	"github.com/kodacampmain/koda-b7-gin/internal/middleware"
	"github.com/kodacampmain/koda-b7-gin/internal/repository"
	"github.com/kodacampmain/koda-b7-gin/internal/service"
)

func RegisterUserRouter(router *gin.Engine, db *pgxpool.Pool) {
	userRouter := router.Group("/users")

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	// userService := service.NewUserServiceMock()
	usersController := controller.NewUsersController(userService)

	userRouter.GET("", usersController.GetAll)
	userRouter.POST("", middleware.MyCustomMiddleware, usersController.Post)
	userRouter.POST("/employee", usersController.Add)
	// params
	userRouter.PUT("/:id/:slug", usersController.Put)

	// protected
	userRouter.GET("/profile", middleware.VerifyToken, usersController.GetProfile)
}
