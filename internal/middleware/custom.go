package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b7-gin/internal/dto"
)

func MyCustomMiddleware(ctx *gin.Context) {
	xkodax := ctx.GetHeader("X-Koda-X")
	// log.Printf("% x\n", xkodax)
	// log.Printf("% x\n", "aku koda")
	if xkodax != "aku koda" {
		ctx.AbortWithStatusJSON(http.StatusConflict, dto.Response{
			Message: "Error",
			Success: false,
			Error:   "wrong usage of header",
		})
		return
	}
	ctx.Next()
}
