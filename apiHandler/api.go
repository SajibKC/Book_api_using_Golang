package apihandler

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Book struct represents a book with its properties.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// In-memory storage for books. Replace with a database for persistence.
var books = []Book{
	{ID: "1", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien"},
	{ID: "2", Title: "Pride and Prejudice", Author: "Jane Austen"},
}

// getBooks handles the retrieval of all books.
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// createBook handles the creation of a new book.
func CreateBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID handles the retrieval of a book by its ID.
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, aBook := range books {
		if aBook.ID == id {
			c.IndentedJSON(http.StatusOK, aBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// updateBookByID handles the update of a book by its ID.
func UpdateBookByID(c *gin.Context) {
    id := c.Param("id")
    var updatedBook Book

    if err := c.BindJSON(&updatedBook); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
        return
    }

    for i, aBook := range books {
        if aBook.ID == id {
            books[i] = updatedBook
            c.IndentedJSON(http.StatusOK, updatedBook)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// deleteBookByID handles the deletion of a book by its ID.
func DeleteBookByID(c *gin.Context) {
    id := c.Param("id")

    for i, aBook := range books {
        if aBook.ID == id {
            // Remove the book from the slice using the index.
            books = append(books[:i], books[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
