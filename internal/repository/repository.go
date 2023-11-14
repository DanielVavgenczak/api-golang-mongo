package repository

import (
	"context"
	"log"

	"github.com/DanielVavgenczak/api-golang-mongo/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryDB struct {
	Mdb *mongo.Database
}

func NewRepositoryDB(db *mongo.Database) *RepositoryDB {
	return &RepositoryDB{
		Mdb: db,
	}
}

func (rb *RepositoryDB) Create(ctx context.Context, movie entity.Movie)(*entity.Movie, error) {
	insert,	err := rb.Mdb.Collection("movies").InsertOne(ctx, &movie)
	if err != nil {
		return nil, err
	}

	// set unique value 
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}
	
	if _, err = rb.Mdb.Collection("movies").Indexes().CreateOne(ctx, index); err != nil {
		return nil, err
	}
	var movieInsert entity.Movie
	query := bson.M{"_id": insert.InsertedID}

	err = rb.Mdb.Collection("movies").FindOne(ctx, query).Decode(&movieInsert)
	if err != nil {
		return nil, err
	}
	log.Println("utimo inserted", movieInsert)
	return &movieInsert, nil

}

func (rb *RepositoryDB) FindByName(ctx context.Context, name string) *entity.Movie{
	var movie entity.Movie
	err := rb.Mdb.Collection("movies").FindOne(ctx, bson.M{"name": name}).Decode(&movie)
	if err != nil {
		return &entity.Movie{}
	}
	return &movie
}