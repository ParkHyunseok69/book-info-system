package main

import (
	"book_info_system/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books/:title", handler.ViewBook)
	router.GET("/books", handler.ListBooks)
	router.POST("/books", handler.AddBook)
	router.PUT("/books/:title", handler.UpdateBook)
	router.DELETE("/books/:title", handler.DeleteBook)

	router.Run("localhost:8080")
}
