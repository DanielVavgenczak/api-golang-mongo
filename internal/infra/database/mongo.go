package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(ctx context.Context, url string)(*mongo.Collection, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	client.Ping(ctx, nil)
	fmt.Print("Connected mongodb")

	collection := client.Database("db_movies").Collection("movies")
	return collection, nil
	
}