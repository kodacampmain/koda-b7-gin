package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kodacampmain/koda-b7-gin/internal/dto"
	"github.com/kodacampmain/koda-b7-gin/pkg"
)

func VerifyToken(ctx *gin.Context) {
	// mengambil token dari request payload (header)
	// header Authorization
	// Bearer token => token diawali dengan kata Bearer
	// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlcm5hbWUiOiJ2YW5kbyIsImlzcyI6ImtvZGEiLCJleHAiOjE3Nzg4MzYzNTl9.grjh7IVfW4FX6MgEDZIEcX-047zyYoPOXvIfCuUMUyE
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Message: "Unauthorized Access, Please Login",
			Success: false,
			Error:   "unauthorized access, please login",
		})
		return
	}
	splittedBearer := strings.Split(bearerToken, " ")
	if len(splittedBearer) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Message: "Unauthorized Access, Please Login",
			Success: false,
			Error:   "invalid token",
		})
		return
	}
	token := splittedBearer[1]

	// verifikasi token nya
	var claims pkg.Claims
	if err := claims.VerifyJWT(token); err != nil {
		log.Println("Error: ", err.Error())
		if errors.Is(err, jwt.ErrTokenInvalidIssuer) || errors.Is(err, jwt.ErrTokenExpired) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Message: "Unauthorized Access, Please Login",
				Success: false,
				Error:   err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Message: "Error",
			Success: false,
			Error:   "Internal Server Error",
		})
		return
	}

	// menempelkan (attach) claims ke context request
	ctx.Set("claims", claims)
	ctx.Next()
}
