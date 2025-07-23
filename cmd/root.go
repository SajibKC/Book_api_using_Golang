package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	apiHandler"example/Book_api_using_Golang/apiHandler"
	"github.com/gin-gonic/gin"
)

var rootCmd = &cobra.Command{
	Use:   "bookapi",
	Short: "Run the Book API server",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.GET("/books", apiHandler.GetBooks)
		r.GET("/books/:id", apiHandler.GetBookByID)
		r.POST("/books", apiHandler.CreateBook)
		r.PUT("/books/:id", apiHandler.UpdateBookByID)
		r.DELETE("/books/:id", apiHandler.DeleteBookByID)

		err := r.Run(":8080")
		if err != nil {
			fmt.Println("Failed to run server:", err)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
