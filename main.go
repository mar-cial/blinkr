package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mar-cial/blinkr/api"
	"github.com/mar-cial/blinkr/db"
	"github.com/mar-cial/blinkr/model"
)

func sep(t string, sym string) {
	maxLength := 40
	wordLen := len(t)
	numberOfDashes := maxLength - wordLen
	rDash := strings.Repeat(sym, numberOfDashes/2)
	fmt.Println(fmt.Sprintf("%s %s %s", rDash, strings.ToUpper(t), rDash))
}

func main() {
	ctx := context.Background()

	client, err := db.CreateClient(os.Getenv("MONGOURI"))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))

	testBlink1 := model.GenerateRandomBlink()
	insertOneRes, err := db.InsertOne(ctx, coll, testBlink1)
	if err != nil {
		log.Fatal(err)
	}

	// title
	sep("main", "=")

	// db actions
	sep("db actions", "-")
	fmt.Println("These are the db action functions.")
	fmt.Println("These are the simplest to set up, there shouldnt be that much of a problem")
	fmt.Println("insert one result")
	fmt.Println(insertOneRes)

	testBlinks := model.GenerateRandomBlinkList(10)
	insertManyRes, err := db.InsertMany(ctx, coll, testBlinks)
	fmt.Println("insert many result")
	fmt.Println(insertManyRes)

	// api
	sep("api routes", "-")
	router, err := api.CreateRouter(os.Getenv("MONGOURI"))
	if err != nil {
		log.Fatal(err)
	}

	port := fmt.Sprintf(":%s", os.Getenv("SERVERPORT"))
	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}
