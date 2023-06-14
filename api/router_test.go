package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRouter(t *testing.T) {
	uri := os.Getenv("MONGOURI")
	c, err := CreateRouter(uri)
	assert.NoError(t, err)
	fmt.Println(c)
}
