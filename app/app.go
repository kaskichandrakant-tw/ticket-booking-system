package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Response struct {
	Success bool `json:"success"`
	Errors []string `json:"errors"`
	Data []Ticket `json:"data"`
}
func StartServer() {
	r := gin.New() // attaches the recovery middleware
	//r.Use(gin.Recovery()) // not sure how to check
	r.GET("/",pingHandler)
	r.POST("/book",handleBookApi )

	r.Run() // lis

	fmt.Println("Server Started at port 8080")
}

func pingHandler (c *gin.Context) {
	c.JSON(200,gin.H{})
}


func handleBookApi(c *gin.Context){
	c.JSON(200,Response{
		Success: true,
		Errors:  []string{},
		Data: []Ticket{
			Ticket{
				Id:      0,
				Catalog: Catalog{
					Id:   0,
					Name: "Movie 1",
				},
				Slot:    Slot{
					Id:   0,
					Date: time.Time{},
				},
			},
		},
	})
}
