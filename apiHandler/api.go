package apihandler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "example/Book_api_using_Golang/dataHandler"
)


func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, datahandler.Books)
}


func CreateBook(c *gin.Context) {
    var newBook datahandler.Book

    if err := c.BindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if newBook.Title == "" || newBook.Author == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Author are required"})
        return
    }

    for _, b := range datahandler.Books {
        if b.ID == newBook.ID {
            c.JSON(http.StatusConflict, gin.H{"error": "Book with this ID already exists"})
            return
        }
    }
    datahandler.Books = append(datahandler.Books, newBook)

    c.IndentedJSON(http.StatusCreated, newBook)
}



func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, aBook := range datahandler.Books {
		if aBook.ID == id {
			c.IndentedJSON(http.StatusOK, aBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}


func UpdateBookByID(c *gin.Context) {
    id := c.Param("id")
    var updatedBook datahandler.Book

    if err := c.BindJSON(&updatedBook); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
        return
    }

    for i, aBook := range datahandler.Books {
        if aBook.ID == id {
            datahandler.Books[i] = updatedBook
            c.IndentedJSON(http.StatusOK, updatedBook)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}


func DeleteBookByID(c *gin.Context) {
    id := c.Param("id")

    for i, aBook := range datahandler.Books {
        if aBook.ID == id {
            datahandler.Books = append(datahandler.Books[:i], datahandler.Books[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func CreateUSer(c *gin.Context) {
    var newUser datahandler.User

    if err := c.BindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if newUser.Username == "" || newUser.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username and Password are required"})
        return
    }

    for _, user := range datahandler.Users {
        if user.Username == newUser.Username {
            c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
            return
        }
    }

    datahandler.Users = append(datahandler.Users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}

