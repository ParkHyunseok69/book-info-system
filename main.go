package main

import (
	"book_info_system/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	router.GET("/book/:id", handler.ViewBook)
	router.GET("/books", handler.ListBooks)
	router.POST("/book", handler.AddBook)
	router.PUT("/book/:id", handler.UpdateBook)
	router.DELETE("/book/:id", handler.DeleteBook)

	router.Run("localhost:8080")
}
