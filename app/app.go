package app

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Response struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
	Data    []Ticket `json:"data"`
}

func StartServer() {
	r := gin.Default() // attaches the recovery middleware
	//r.Use(gin.Recovery()) // not sure how to check
	v1 := r.Group("/v1")
	{
		v1.GET("/", v1Handler)
		v1.POST("/book", handleBookApi)
	}

	r.GET("/", pingHandler)
	r.Run()
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message":"welcome to app"})
}

func v1Handler(c *gin.Context) {
	c.JSON(200, gin.H{"message":"welcome to v1"})
}

func handleBookApi(c *gin.Context) {
	tickets := []Ticket{
		{
			Id: 0,
			Catalog: Catalog{
				Id:   0,
				Name: "Movie 1",
			},
			Slot: Slot{
				Id:   0,
				Date: time.Time{},
			},
		},
	}
	c.JSON(200, Response{
		Success: true,
		Errors:  []string{},
		Data:    tickets,
	})
}
