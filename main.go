package main

import (
	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/db"
	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", routes.GetEvents)
	server.POST("/events", routes.CreateEvent)
	server.GET("/events/:id", routes.GetEventById)

	server.Run(":3333")
}
