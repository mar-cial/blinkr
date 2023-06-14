package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/mar-cial/blinkr/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {
	client, err := CreateClient(os.Getenv("MONGOURI"))
	assert.NoError(t, err)
	assert.NotEmpty(t, client)
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	if err := utils.LoadTestEnvs(); err != nil {
		fmt.Println("Failed to load envs")
		log.Fatal(err)
	}

	mongoC, err := utils.CreateMongoContainer()
	if err != nil {
		fmt.Println("Failed to create Mongo container")
		log.Fatal(err)
	}

	m.Run()

	if err := utils.TerminateMongoContainer(ctx, mongoC); err != nil {
		fmt.Println("Failed to terminate Mongo container")
		log.Fatal(err)
	}
}
