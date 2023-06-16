package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mar-cial/blinkr/api"
	"github.com/mar-cial/blinkr/db"
)

func main() {
	ctx := context.Background()

	client, err := db.CreateClient(os.Getenv("MONGOURI"))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	router, err := api.CreateRouter(os.Getenv("MONGOURI"))
	if err != nil {
		log.Fatal(err)
	}

	port := fmt.Sprintf(":%s", os.Getenv("SERVERPORT"))
	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}
