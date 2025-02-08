package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct{
	collection *mongo.Collection
	ctx context.Context
}

func New() (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/main")
	client, err := mongo.NewClient(clientOptions)
	if err != nil{
		return nil, err
	}
	coll := client.Database("main").Collection("arts")
	return &Storage{
		collection: coll,
		ctx: ctx,
	}, nil
}
