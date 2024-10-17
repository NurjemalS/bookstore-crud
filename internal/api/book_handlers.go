package api

import (
	"bookstore/internal/models"
	repository "bookstore/internal/store"
	"context"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func PostBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newBook.FilePath = ""
	err := repository.AddBook(context.Background(), newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBooks(c *gin.Context) {
	books, err := repository.GetAllBooks(context.Background())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := repository.GetBookByID(context.Background(), id)
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

	err := repository.UpdateBook(context.Background(), id, updatedBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteBook(context.Background(), id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func UploadFileForBook(c *gin.Context) {

	bookID := c.Param("id")

	// Get file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	destination := filepath.Join("uploads", file.Filename)
	log.Printf(destination)

	// Save file
	if err := c.SaveUploadedFile(file, destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Update
	if err := repository.UpdateBookFilePath(context.Background(), bookID, destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book with file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded for a book", "file_path": destination})
}
