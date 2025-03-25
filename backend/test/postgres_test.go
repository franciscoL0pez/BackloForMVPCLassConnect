package test

import (
	"MVPCLassConnect/backend/dbConfig/postgres"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load("../../postgres.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestConnectDB(t *testing.T) {

	loadEnv()
	postgres.ConnectDB()

	if postgres.DB == nil {
		t.Fatalf("Expected DB connection, but got nil")
	}

	postgres.CloseDB()
}
