package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// format of book
type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// slice storage of books
var books = []book{
	{ID: "1", Title: "Mahabharata", Author: "VedVyasa", Price: 3000},
	{ID: "2", Title: "Ramayana", Author: "Valmiki", Price: 2700},
}

func main() {
	fmt.Println("Book Godown is open:")

	router := gin.Default()

	// routers for all CRUD requests
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", addBook)

	router.Run("localhost:8989")
}

// to get all books from godown
func getBooks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, books)
}

// to get one book by id from godown
func getBookById(c *gin.Context) {

	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// to add one book to godown
func addBook(c *gin.Context) {

	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// adding newBook to slice storage
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, books)
}
