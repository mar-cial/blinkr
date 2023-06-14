package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/mar-cial/blinkr/model"
	"github.com/stretchr/testify/assert"
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
