package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quality int `json:"quality"`
}

var books = []book{
	{ID: "1", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quality: 9},
	{ID: "2", Title: "Cloud Native Go", Author: "M.-L. Reimer", Quality: 9},
	{ID: "3", Title: "The Hobbit", Author: "J. R. R. Tolkien", Quality: 7},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(books); i++ {
		if books[i].ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.DELETE("/books/:id", deleteBook)
	router.GET("/", home)
	router.Run("localhost:8080")
}
