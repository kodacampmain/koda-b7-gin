package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b7-gin/internal/router"
)

func main() {
	// inisialisasi
	// gin.New()
	app := gin.Default()
	// install router
	router.InitRouter(app)
	// run
	app.Run("localhost:8080")
}
