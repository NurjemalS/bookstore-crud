package api

import (
	"context"
	"net/http"
	"bookstore/internal/models"
	"bookstore/internal/store"
	"github.com/gin-gonic/gin"
)

func PostBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := bookRepository.AddBook(context.Background(), newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBooks(c *gin.Context) {
	books, err := bookRepository.GetAllBooks(context.Background())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := bookRepository.GetBookByID(context.Background(), id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book

	if err := c.BindJSON(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := bookRepository.UpdateBook(context.Background(), id, updatedBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := bookRepository.DeleteBook(context.Background(), id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
