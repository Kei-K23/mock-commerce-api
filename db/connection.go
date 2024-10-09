package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDB() {
	dsn := os.Getenv("DB_CONNECTION_URL")
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Error when parsing db config: %v", err)
	}

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Error when connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to database")
}
