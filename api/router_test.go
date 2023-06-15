package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRouter(t *testing.T) {
	uri := os.Getenv("MONGOURI")
	assert.NotEmpty(t, uri)

	c, err := CreateRouter(uri)
	assert.NoError(t, err)
	assert.NotEmpty(t, c)
}
