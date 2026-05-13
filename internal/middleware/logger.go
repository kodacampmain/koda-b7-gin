package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(ctx *gin.Context) {
	start := time.Now()
	date := fmt.Sprintf("%d/%d/%d", start.Day(), start.Month(), start.Year())
	clock := fmt.Sprintf("%d:%d:%d", start.Hour()+1, start.Minute(), start.Second())
	origin := ctx.GetHeader("Origin")
	method := ctx.Request.Method
	url := ctx.Request.URL
	// happen before controller
	ctx.Next()
	// happen after response sent by controller
	end := time.Now()
	duration := end.Sub(start)
	status := ctx.Writer.Status()
	log.Printf("[LOGGER] %s - %s | %s | %s | %s | %dms | %d", date, clock, origin, url.Path, method, duration/time.Millisecond, status)
}
