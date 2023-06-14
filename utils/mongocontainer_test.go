package utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

var (
	c testcontainers.Container
)

func TestCreateMongoContainer(t *testing.T) {
	mongoC, err := CreateMongoContainer()
	assert.NoError(t, err)
	assert.NotEmpty(t, mongoC.GetContainerID())
	assert.Len(t, mongoC.GetContainerID(), 64)

	c = mongoC
}

func TestTerminateMongoContainer(t *testing.T) {
	ctx := context.Background()
	err := TerminateMongoContainer(ctx, c)
	assert.NoError(t, err)
}
