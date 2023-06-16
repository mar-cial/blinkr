package api

import (
	"fmt"
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

func (b *blinksController) insertMany(c *gin.Context) {
	coll := b.mc.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))
	var input []model.Blink
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := db.InsertMany(c, coll, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (b *blinksController) listOne(c *gin.Context) {
	coll := b.mc.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))
	id := c.Param("id")

	res, err := db.ListOne(c, coll, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (b *blinksController) listAll(c *gin.Context) {
	coll := b.mc.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))

	res, err := db.ListAll(c, coll)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (b *blinksController) updateOne(c *gin.Context) {
	coll := b.mc.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))
	id := c.Param("id")

	var input model.Blink
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println()
	fmt.Println("update one controller input")
	fmt.Println(input)
	fmt.Println()

	fmt.Println("id sent")
	fmt.Println(id)

	res, err := db.UpdateOne(c, coll, id, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (b *blinksController) deleteOne(c *gin.Context) {
	coll := b.mc.Database(os.Getenv("DBNAME")).Collection(os.Getenv("DBCOLL"))
	id := c.Param("id")

	res, err := db.DeleteOne(c, coll, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
