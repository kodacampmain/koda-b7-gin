package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kodacampmain/koda-b7-gin/internal/dto"
)

type MoviesController struct{}

func NewMoviesController() *MoviesController {
	return &MoviesController{}
}

func (m *MoviesController) Search(ctx *gin.Context) {
	var mq dto.MoviesQuery
	if err := ctx.ShouldBindWith(&mq, binding.Query); err != nil {
		log.Println("Error: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Message: "Error",
			Data:    nil,
			Success: false,
			Error:   "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Message: "OK",
		Data:    mq,
		Success: true,
		Error:   "",
	})
}
