package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b7-gin/pkg"
)

func Rbac(ctx *gin.Context) {
	// cek role
	token, _ := ctx.Get("claims")

	// type assertion
	claims := token.(pkg.Claims)
	log.Println(claims)
	// kalau tidak boleh, abort dengan 403
	// if claims.Role {}
	// kalau boleh, maka next
}
