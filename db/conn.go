package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
	"time"
)

var connUrl = os.Getenv("ME_CONFIG_MONGODB_URL")
var dbName = os.Getenv("DB_NAME")
var colName = os.Getenv("COL_NAME")

var client *mongo.Client
var once sync.Once

func getClient() *mongo.Client {
	var clientError error
	once.Do(func() {
		client, clientError = mongo.NewClient(options.Client().ApplyURI(connUrl))
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := client.Connect(ctx)
		if err != nil {
			clientError = err
		}
		err = client.Ping(ctx, nil)
		if err != nil {
			clientError = err
		}
		cancel()
	})
	if clientError != nil {
		log.Fatal(clientError)
	}
	return client
}
