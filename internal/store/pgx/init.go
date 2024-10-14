package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"fmt"
)

var dbpool *pgxpool.Pool

func InitDB(username, password, host, port, dbName string) (*pgxpool.Pool, error) {
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
	var err error
	dbpool, err = pgxpool.New(context.Background(), connectionStr)
	if err != nil {
		log.Printf("Unable to connect to database: %v", err)
		return nil, err
	}
	return dbpool, nil
}

func GetDBPool() *pgxpool.Pool {
	return dbpool
}
