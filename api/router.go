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
	r.POST("/blinks/create/many", bc.insertMany)
	r.GET("/blinks/list/:id", bc.listOne)
	r.GET("/blinks/list", bc.listAll)
	r.PUT("/blinks/update/:id", bc.updateOne)
	r.DELETE("/blinks/delete/:id", bc.deleteOne)

	return r, nil
}
