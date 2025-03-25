package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() {
	
	err := godotenv.Load("../../postgres.env")
	if err != nil {
		log.Println("Not possible load .env file")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Not its possible connect: %v", err)
	}

	
	DB = pool
	fmt.Println("Connected to PostgreSQL")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Close connection to PostgreSQL")
	}
}
