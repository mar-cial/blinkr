package model

import (
	"encoding/json"

	lorem "github.com/drhodes/golorem"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UnmarshalBlink(data []byte) (Blink, error) {
	var r Blink
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Blink) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func GenerateRandomBlink() Blink {
	t := lorem.Sentence(3, 8)
	m := lorem.Paragraph(3, 5)

	return Blink{
		Title:   t,
		Message: m,
	}
}

func GenerateRandomBlinkList(qty int) []Blink {
	var blinks []Blink

	for a := 0; a < qty; a++ {
		blink := GenerateRandomBlink()
		blinks = append(blinks, blink)
	}

	return blinks
}

type Blink struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title   string             `json:"title"`
	Message string             `json:"message"`
}
