package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func getCollection() *mongo.Collection {
	client := getClient()
	return client.Database(dbName).Collection(colName)
}

func Get(f bson.M, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return getCollection().FindOne(ctx, f, opts...)
}

func InsertDoc(doc any, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	one, err := getCollection().InsertOne(ctx, doc, opts...)
	return one, err
}

func UpdateDocs(f bson.D, u bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return getCollection().UpdateMany(ctx, f, u, opts...)
}

func UpdateDoc(f bson.D, u bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return getCollection().UpdateOne(ctx, f, u, opts...)
}
