package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kodacampmain/koda-b7-gin/internal/dto"
)

type IUserService interface {
	PrintUser(body dto.UsersBody)
	ValidateEmail(email string) error
}

type UsersController struct {
	userService IUserService
}

func NewUsersController(userService IUserService) *UsersController {
	return &UsersController{
		userService: userService,
	}
}

func (u *UsersController) Post(ctx *gin.Context) {
	var body dto.UsersBody
	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		// apakah error validasi?
		if strings.Contains(err.Error(), "Email") {
			if strings.Contains(err.Error(), "required") {
				log.Println("Error: ", "Missing Email")
				ctx.JSON(http.StatusBadRequest, dto.Response{
					Message: "Bad Request",
					Success: false,
					Error:   "email is required",
				})
				return
			}
			if strings.Contains(err.Error(), "email") {
				log.Println("Error: ", "Invalid Email Format")
				ctx.JSON(http.StatusBadRequest, dto.Response{
					Message: "Bad Request",
					Success: false,
					Error:   "email is using invalid format",
				})
				return
			}
		}
		// kirim response error
		log.Println("Error: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Message: "Error",
			Data:    nil,
			Success: false,
			Error:   "Internal Server Error",
		})
		return
	}
	err := u.userService.ValidateEmail(body.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Message: "Invalid Email Format",
			Error:   err.Error(),
			Success: false,
		})
		return
	}

	u.userService.PrintUser(body)
	// kirim response sukses
	ctx.JSON(http.StatusOK, dto.Response{
		Message: "OK",
		Data:    body,
		Success: true,
		Error:   "",
	})
}

func (u *UsersController) Put(ctx *gin.Context) {
	var uri dto.UsersUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
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
		Data:    uri,
		Success: true,
		Error:   "",
	})
}
