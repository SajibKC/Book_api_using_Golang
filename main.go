package main

import (
	apiHandler "example/Book_api_using_Golang/apiHandler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define routes for book operations.
	router.GET("/books", apiHandler.GetBooks)
	router.GET("/books/:id", apiHandler.GetBookByID)
	router.POST("/books", apiHandler.CreateBook)
	router.PUT("/books/:id", apiHandler.UpdateBookByID)
	router.DELETE("/books/:id", apiHandler.DeleteBookByID)

	router.Run("localhost:8080")
}