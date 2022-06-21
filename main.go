package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var Books = []Book{
	{Id: "1", Title: "sailing sun set", Author: "Khali", Quantity: 3},
	{Id: "2", Title: "Docking sun set", Author: "Joe", Quantity: 6},
	{Id: "3", Title: "Seting sun set", Author: "Mukundi", Quantity: 15},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Books)
}
func getBook(c *gin.Context) {
	// to geT the dynamic parameter we do this
	id := c.Param("id")
	//NOTE  for query parameters we do this
	//  id, err := c.GetQuery("id")

	book, err := getBookById(id)
	if err != nil {
		// gin.H helps us write json or rather helsp us just create ajson format of this type quickly

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
func getBookById(id string) (*Book, error) {
	for i, b := range Books {
		if b.Id == id {
			return &Books[i], nil
		}
	}
	return nil, errors.New("no book was found with that id")
}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	Books = append(Books, newBook)
	c.IndentedJSON(http.StatusCreated, Books)

}

func main() {

	// set up router with gin just like route in node js
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("book/:id", getBook)

	router.Run("localhost:8080")

	// fmt.Println("hellow world", Books)
}
