package api

import (
	"github.com/gin-gonic/gin"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
	}
}

func CreateRouter(uri string) (*gin.Engine, error) {
	bc, err := CreateBlinksController(uri)
	if err != nil {
		return nil, err
	}

	r := gin.Default()

	// middleware
	r.Use(middleware())

	// routes
	r.POST("/blinks/create/one", bc.insertOne)

	return r, nil
}
