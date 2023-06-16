package db

import (
	"context"
	"fmt"

	"github.com/mar-cial/blinkr/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	fmt.Println("received blinks in insert many db action: ")
	for b := range in {
		fmt.Printf("title: \t%s\n", in[b].Title)
		fmt.Printf("message: \t%s\n", in[b].Message)
	}

	return coll.InsertMany(ctx, docs, nil)
}

func ListOne(ctx context.Context, coll *mongo.Collection, id string) (model.Blink, error) {
	var blink model.Blink
	mongoid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return blink, err

	}

	if err := coll.FindOne(ctx, bson.M{"_id": mongoid}, nil).Decode(&blink); err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("no documents found")
			return blink, err
		}
		return blink, err
	}

	return blink, err
}

func ListAll(ctx context.Context, coll *mongo.Collection) ([]model.Blink, error) {
	var blinks []model.Blink

	cur, err := coll.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("err finding docs")
		fmt.Println(err.Error())
		return blinks, err
	}

	if err := cur.All(ctx, &blinks); err != nil {
		fmt.Println("err decoding documents")
		fmt.Println(err.Error())
		return blinks, err
	}

	return blinks, err
}

func UpdateOne(ctx context.Context, coll *mongo.Collection, id string, in model.Blink) (*mongo.UpdateResult, error) {
	var updateResult *mongo.UpdateResult

	mongoid, err := primitive.ObjectIDFromHex(id)
	fmt.Println()
	fmt.Println("MONGOID")
	fmt.Println(mongoid)

	if err != nil {
		return updateResult, err
	}

	doc := bson.M{"title": in.Title, "message": in.Message}
	update := bson.M{"$set": doc}

	return coll.UpdateByID(ctx, mongoid, update)
}

func DeleteOne(ctx context.Context, coll *mongo.Collection, id string) (*mongo.DeleteResult, error) {
	var deleteRes *mongo.DeleteResult

	mongoid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return deleteRes, err
	}

	return coll.DeleteOne(ctx, bson.M{"_id": mongoid})

}
