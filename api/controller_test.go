package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mar-cial/blinkr/model"
	"github.com/stretchr/testify/assert"
)

var (
	c *blinksController
)

func TestCreateBlinkrController(t *testing.T) {
	controller, err := CreateBlinksController(os.Getenv("MONGOURI"))
	assert.NoError(t, err)

	c = controller
}

func TestInsertOne(t *testing.T) {
	path := "/blinks/create/one"
	router, err := CreateRouter(os.Getenv("MONGOURI"))
	assert.NoError(t, err)

	blink := model.GenerateRandomBlink()

	// blink bytes
	bb, err := blink.Marshal()
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader(bb))
	assert.NoError(t, err)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())
}
