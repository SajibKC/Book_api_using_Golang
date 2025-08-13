package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	apiHandler"example/Book_api_using_Golang/apiHandler"
	"github.com/gin-gonic/gin"
	authHandler "example/Book_api_using_Golang/authHandler"
)

var rootCmd = &cobra.Command{
	Use:   "bookapi",
	Short: "Run the Book API server",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.GET("/books", authHandler.AuthMiddleware(), apiHandler.GetBooks)
		r.GET("/books/:id", authHandler.AuthMiddleware(), apiHandler.GetBookByID)
		r.POST("/books", authHandler.AuthMiddleware(), apiHandler.CreateBook)
		r.PUT("/books/:id", authHandler.AuthMiddleware(), apiHandler.UpdateBookByID)
		r.DELETE("/books/:id", authHandler.AuthMiddleware(), apiHandler.DeleteBookByID)
		r.POST("/users", apiHandler.CreateUSer)

		err := r.Run(":8080")
		if err != nil {
			fmt.Println("Failed to run server:", err)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
