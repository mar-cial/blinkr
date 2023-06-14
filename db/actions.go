package db

import (
	"context"
	"fmt"

	"github.com/mar-cial/blinkr/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(ctx context.Context, coll *mongo.Collection, in model.Blink) (*mongo.InsertOneResult, error) {
	var doc interface{}
	doc = in

	fmt.Println("Received blink: ", in)
	fmt.Println("Assigned doc: ", doc)

	return coll.InsertOne(ctx, doc)
}

func InsertMany(ctx context.Context, coll *mongo.Collection, in []model.Blink) (*mongo.InsertManyResult, error) {
	var docs []interface{}

	for a := range in {
		docs = append(docs, in[a])
	}

	return coll.InsertMany(ctx, docs, nil)
}
