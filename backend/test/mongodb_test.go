// example test
package test

import (
	"MVPCLassConnect/backend/dbConfig/mongo"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMongoDBConnection(t *testing.T) {
	// Attempt to connect to MongoDB
	mongo.ConnectDB()

	// Ensure MongoClient is initialized
	assert.NotNil(t, mongo.MongoClient, "MongoDB client should not be nil after connection")

	// Check if the database is reachable using Ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := mongo.MongoClient.Ping(ctx, nil)
	assert.NoError(t, err, "Ping to MongoDB should not return an error")

	// Close the database connection
	mongo.CloseDB()
	assert.Nil(t, mongo.MongoClient, "MongoClient should be nil after closing the connection")
}
