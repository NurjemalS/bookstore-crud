package main

import (
	"bookstore/config"
	"bookstore/internal/store/pgx"
	"bookstore/internal/api"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	cfg := config.LoadConfig()
	dbpool, err := pgx.InitDB(cfg.DbUsername, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbDatabase)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	router := gin.Default()
	router.GET("/books", api.GetBooks)
	router.POST("/books", api.PostBook)
	router.GET("/books/:id", api.GetBookById)
	router.PUT("/books/:id", api.UpdateBook)
	router.DELETE("/books/:id", api.DeleteBook)

	// Start the server
	// router.Run(cfg.ServerAddress)
	router.Run("localhost:8080")
}
