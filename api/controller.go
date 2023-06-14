package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mar-cial/blinkr/db"
	"github.com/mar-cial/blinkr/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type blinksController struct {
	mc *mongo.Client
}

func CreateBlinksController(uri string) (*blinksController, error) {
	c, err := db.CreateClient(uri)
	if err != nil {
		return nil, err
	}

	return &blinksController{
		mc: c,
	}, nil
}

func (b *blinksController) insertOne(c *gin.Context) {
	coll := b.mc.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))
	var input model.Blink
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := db.InsertOne(c, coll, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
