package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func createCollection(opts ...*options.CreateCollectionOptions) error {
	c := getClient()
	err := c.Database(dbName).CreateCollection(context.TODO(), colName, opts...)
	return err
}

func Migrate() {
	if connUrl == "" || dbName == "" || colName == "" {
		log.Fatal("Service need specified 'ME_CONFIG_MONGODB_URL', 'dbName' and 'collectionName' in .env")
	}
	err := createCollection()
	if err != nil {
		switch err.(type) {
		case mongo.CommandError:
			log.Printf("Database '%s' and collection '%s' already exists", dbName, colName)
		default:
			log.Fatal(err)
		}
	}
}
