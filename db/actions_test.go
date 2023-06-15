package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/mar-cial/blinkr/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	testid string
)

func TestInsertOne(t *testing.T) {
	ctx := context.Background()

	uri := os.Getenv("MONGOURI")
	fmt.Println("uri")
	fmt.Println(uri)

	blink := model.GenerateRandomBlink()

	client, err := CreateClient(uri)
	assert.NoError(t, err)

	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))

	res, err := InsertOne(ctx, coll, blink)
	assert.NoError(t, err)

	fmt.Println("test response")
	fmt.Println(res)

	mongoid := res.InsertedID.(primitive.ObjectID)
	testid = mongoid.Hex()
}

func TestInsertMany(t *testing.T) {
	ctx := context.Background()

	uri := os.Getenv("MONGOURI")
	fmt.Println("uri must be set still")
	fmt.Println(uri)

	blinks := model.GenerateRandomBlinkList(5)

	client, err := CreateClient(uri)
	assert.NoError(t, err)

	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))

	res, err := InsertMany(ctx, coll, blinks)
	assert.NoError(t, err)

	fmt.Println("test response")
	fmt.Println(res)
}

func TestListOne(t *testing.T) {
	ctx := context.Background()

	uri := os.Getenv("MONGOURI")
	fmt.Println("uri must be set still")
	fmt.Println(uri)

	client, err := CreateClient(uri)
	assert.NoError(t, err)

	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))

	blinkRes, err := ListOne(ctx, coll, testid)
	assert.NoError(t, err)

	assert.True(t, primitive.IsValidObjectID(blinkRes.ID.Hex()))
	assert.NotEmpty(t, blinkRes.Title)
	assert.NotEmpty(t, blinkRes.Message)
	fmt.Println("test list one response")
	fmt.Println(blinkRes)
}

func TestListAll(t *testing.T) {
	ctx := context.Background()

	uri := os.Getenv("MONGOURI")
	fmt.Println("uri must be set still")
	fmt.Println(uri)

	client, err := CreateClient(uri)
	assert.NoError(t, err)

	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))

	blinkRes, err := ListAll(ctx, coll)
	assert.NoError(t, err)

	for a := range blinkRes {
		b := blinkRes[a]

		assert.True(t, primitive.IsValidObjectID(b.ID.Hex()))
		assert.NotEmpty(t, b.Title)
		assert.NotEmpty(t, b.Message)
		fmt.Println("test list one response")
		fmt.Println(b)
	}
}

func TestUpdateOne(t *testing.T) {
	ctx := context.Background()

	uri := os.Getenv("MONGOURI")
	fmt.Println("uri must be set still")
	fmt.Println(uri)

	client, err := CreateClient(uri)
	assert.NoError(t, err)

	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))

	in := model.Blink{
		Title:   "This is an updated blink",
		Message: "This is the message of an updated blink",
	}

	updateRes, err := UpdateOne(ctx, coll, testid, in)
	assert.NoError(t, err)

	fmt.Println("Update result")
	fmt.Println(updateRes)
}

func TestDeleteOne(t *testing.T) {
	ctx := context.Background()

	uri := os.Getenv("MONGOURI")
	fmt.Println("uri must be set still")
	fmt.Println(uri)

	client, err := CreateClient(uri)
	assert.NoError(t, err)

	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))

	delRes, err := DeleteOne(ctx, coll, testid)
	assert.NoError(t, err)

	fmt.Println("Delete result")
	fmt.Println(delRes)
}
