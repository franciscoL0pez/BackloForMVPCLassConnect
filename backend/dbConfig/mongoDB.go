// create a connect to mongo db

package dbConfig

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// this uri we need to move to .env file
const mongoURI = "mongodb://localhost:27017"

// MongoClient instance db
var MongoClient *mongo.Client

func ConnectDB() {
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Conectar con MongoDB
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error to create the client %v", err)
	}

	// Definir contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error to connect db %v", err)
	}

	// ping to verify the conecction
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	fmt.Println("Conectado a MongoDB")
	MongoClient = client
}


func CloseDB() {
	if MongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Disconnect ping to verify the connection
		err := MongoClient.Disconnect(ctx)
		if err != nil {
			log.Fatalf("Error closing connection to MongoDB: %v", err)
		}
		fmt.Println("Conexión close")
	
		MongoClient = nil
	}
}

