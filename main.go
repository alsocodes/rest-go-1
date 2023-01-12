package main

import (
	"log"
	"net/http"
	bookController "rest/controllers"
	"rest/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.ConnectDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})

	r.GET("/books", bookController.FindBooks)
	r.GET("/books/:id", bookController.FindBook)
	r.POST("/books", bookController.CreateBook)
	r.PUT("/books/:id", bookController.UpdateBook)
	r.DELETE("/books/:id", bookController.DeleteBook)

	err := r.Run()
	if err != nil {
		return
	}
}
