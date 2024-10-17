package pgx

import (
	errors "bookstore/internal/app"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbpool *pgxpool.Pool

func InitDB(username, password, host, port, dbName string) (*pgxpool.Pool, error) {
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)

	dbpool, err := pgxpool.New(context.Background(), connectionStr)
	if err != nil {
		log.Printf("Unable to connect to database: %v", err)

		return nil, errors.NewCustomError("Failed to connect to database", 500)
	}

	return dbpool, nil
}
func GetDBPool() *pgxpool.Pool {
	return dbpool
}
