package api

import (
	"bookstore/internal/models"
	repository "bookstore/internal/store"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAuthor(c *gin.Context) {
	var newAuthor models.Author

	if err := c.BindJSON(&newAuthor); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := repository.AddAuthor(context.Background(), newAuthor)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add author"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAuthor)
}

func GetAuthors(c *gin.Context) {
	authors, err := repository.GetAllAuthor(context.Background())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch authors"})
		return
	}

	c.IndentedJSON(http.StatusOK, authors)
}

func GetAuthorById(c *gin.Context) {
	id := c.Param("id")
	author, err := repository.GetAuthorByID(context.Background(), id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, author)
}

func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var updatedAuthor models.Author

	if err := c.BindJSON(&updatedAuthor); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := repository.UpdateAuthor(context.Background(), id, updatedAuthor)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Author"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedAuthor)
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteAuthor(context.Background(), id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
