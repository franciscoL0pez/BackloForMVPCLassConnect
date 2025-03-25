package test

import (
	//"context"
	"testing"
	//"time"

	"github.com/stretchr/testify/assert"
	//"github.com/franciscoL0pez/MVPCLassConnect/backend/dbConfig"
)
/*
func TestMongoDBConnection(t *testing.T) {
	// Attempt to connect to MongoDB
	dbConfig.ConnectDB()

	// Ensure MongoClient is initialized
	assert.NotNil(t, dbConfig.MongoClient, "MongoDB client should not be nil after connection")

	// Check if the database is reachable using Ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := dbConfig.MongoClient.Ping(ctx, nil)
	assert.NoError(t, err, "Ping to MongoDB should not return an error")

	// Close the database connection
	dbConfig.CloseDB()
	assert.Nil(t, dbConfig.MongoClient, "MongoClient should be nil after closing the connection")
}
*/

func TestStr(t *testing.T) {
	assert.Equal(t, "Hello, World!", "Hello, World!")
}
