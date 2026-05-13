package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b7-gin/internal/controller"
	"github.com/kodacampmain/koda-b7-gin/internal/service"
)

func RegisterUserRouter(router *gin.Engine) {
	userRouter := router.Group("/users")

	userService := service.NewUserService()
	// userService := service.NewUserServiceMock()
	usersController := controller.NewUsersController(userService)

	userRouter.POST("", usersController.Post)
	// params
	userRouter.PUT("/:id/:slug", usersController.Put)
}
