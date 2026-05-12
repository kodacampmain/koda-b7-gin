package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	// inisialisasi
	// gin.New()
	router := gin.Default()
	// implementasi route
	// router.METHOD(endpoint, callback)
	router.GET("/", func(c *gin.Context) {
		// mengirim response
		c.JSON(http.StatusOK, Response{
			Message: "hello",
			Data: gin.H{
				"name": "koda",
			},
			Success: true,
			Error:   "",
		})
		// pastikan setelah response tidak ada proses logika
	})
	router.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "hello")
	})
	router.POST("/users", func(ctx *gin.Context) {
		var body UsersBody
		if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
			// kirim response error
			log.Println("Error: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "Error",
				Data:    nil,
				Success: false,
				Error:   "Internal Server Error",
			})
			return
		}
		// proses logika i.e. validasi
		log.Printf("\nNama: %s\nEmail: %s\nAge: %d\nDob: %s\n", body.Fullname, body.Email, body.Age, body.Dob.String())
		// kirim response sukses
		ctx.JSON(http.StatusOK, Response{
			Message: "OK",
			Data:    body,
			Success: true,
			Error:   "",
		})
	})
	router.PUT("/users/:id/:slug", func(ctx *gin.Context) {
		var uri UsersUri
		if err := ctx.ShouldBindUri(&uri); err != nil {
			log.Println("Error: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "Error",
				Data:    nil,
				Success: false,
				Error:   "Internal Server Error",
			})
			return
		}
		ctx.JSON(http.StatusOK, Response{
			Message: "OK",
			Data:    uri,
			Success: true,
			Error:   "",
		})
	})
	router.GET("/movies", func(ctx *gin.Context) {
		var mq MoviesQuery
		if err := ctx.ShouldBindWith(&mq, binding.Query); err != nil {
			log.Println("Error: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "Error",
				Data:    nil,
				Success: false,
				Error:   "Internal Server Error",
			})
			return
		}
		ctx.JSON(http.StatusOK, Response{
			Message: "OK",
			Data:    mq,
			Success: true,
			Error:   "",
		})
	})
	// run
	router.Run("localhost:8080")
}

type UsersBody struct {
	// key datatype `tag`
	Fullname string    `json:"nama_lengkap" form:"nl"`
	Email    string    `json:"surel"`
	Age      int       `json:"umur"`
	Dob      time.Time `json:"ttl"`
}

type UsersUri struct {
	Id   int    `uri:"id" json:"id"`
	Slug string `uri:"slug" json:"slug"`
}

type MoviesQuery struct {
	Title string `form:"title" json:"title"`
	Genre string `form:"genre" json:"genre"`
}

type Response struct {
	Message string
	Data    any
	Success bool
	Error   string
}
