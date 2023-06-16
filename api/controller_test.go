package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mar-cial/blinkr/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	c      *blinksController
	testid string
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

	var insertOneRes *mongo.InsertOneResult
	err = json.NewDecoder(w.Body).Decode(&insertOneRes)
	assert.NoError(t, err)

	testid = insertOneRes.InsertedID.(string)
}

func TestInsertMany(t *testing.T) {
	path := "/blinks/create/many"
	router, err := CreateRouter(os.Getenv("MONGOURI"))
	assert.NoError(t, err)

	blinks := model.GenerateRandomBlinkList(5)

	bbs, err := json.Marshal(blinks)
	assert.NoError(t, err)
	assert.NotEmpty(t, bbs)

	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader(bbs))
	assert.NoError(t, err)

	router.ServeHTTP(w, req)

	// Blinks response
	var insertManyRes *mongo.InsertManyResult
	err = json.NewDecoder(w.Body).Decode(&insertManyRes)
	assert.NoError(t, err)

	for a := range insertManyRes.InsertedIDs {
		b := insertManyRes.InsertedIDs[a].(string)
		assert.NotEmpty(t, b)
		assert.True(t, primitive.IsValidObjectID(b))
	}
}

func TestListOne(t *testing.T) {
	path := fmt.Sprintf("/blinks/list/%s", testid)

	router, err := CreateRouter(os.Getenv("MONGOURI"))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, path, nil)
	assert.NoError(t, err)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var b model.Blink
	err = json.NewDecoder(w.Body).Decode(&b)
	assert.NoError(t, err)

	assert.True(t, primitive.IsValidObjectID(b.ID.Hex()))
	assert.NotEmpty(t, b.Title)
	assert.NotEmpty(t, b.Message)
}

func TestListAll(t *testing.T) {
	path := "/blinks/list"

	router, err := CreateRouter(os.Getenv("MONGOURI"))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, path, nil)
	assert.NoError(t, err)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var blinks []model.Blink
	err = json.NewDecoder(w.Body).Decode(&blinks)
	assert.NoError(t, err)

	for a := range blinks {
		b := blinks[a]

		assert.True(t, primitive.IsValidObjectID(b.ID.Hex()))
		assert.NotEmpty(t, b.Title)
		assert.NotEmpty(t, b.Message)
	}
}

func TestUpdateOne(t *testing.T) {
	path := fmt.Sprintf("/blinks/update/%s", testid)

	router, err := CreateRouter(os.Getenv("MONGOURI"))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	updatedBlink := model.GenerateRandomBlink()

	// updated blink bytes
	ubb, err := updatedBlink.Marshal()
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, path, bytes.NewReader(ubb))
	assert.NoError(t, err)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteOne(t *testing.T) {
	path := fmt.Sprintf("/blinks/delete/%s", testid)

	router, err := CreateRouter(os.Getenv("MONGOURI"))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	assert.NoError(t, err)

	router.ServeHTTP(w, req)
}
