package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var blinkbytes []byte

func TestMarshalBlink(t *testing.T) {
	b := Blink{
		ID:      primitive.NewObjectID(),
		Title:   "Test blink 1",
		Message: "This is the message for test blink 1",
	}

	bbytes, err := b.Marshal()
	assert.NoError(t, err)

	blinkbytes = bbytes

	strBlink := string(bbytes)

	assert.Contains(t, strBlink, b.Title)
	assert.Contains(t, strBlink, b.Message)
}

func TestUnmarshalBlink(t *testing.T) {
	blink, err := UnmarshalBlink(blinkbytes)
	assert.NoError(t, err)

	assert.Equal(t, "Test blink 1", blink.Title)
	assert.Equal(t, "This is the message for test blink 1", blink.Message)
}

func TestGenerateBlink(t *testing.T) {
	b := GenerateRandomBlink()
	assert.NotEmpty(t, b.Title)
	assert.NotEmpty(t, b.Message)
}

func TestGenerateBlinkList(t *testing.T) {
	bs := GenerateRandomBlinkList(10)
	assert.Len(t, bs, 10)

	for a := range bs {
		b := bs[a]

		assert.NotEmpty(t, b.Title)
		assert.NotEmpty(t, b.Message)
	}
}
