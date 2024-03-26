package main

import (
	"net/http"

	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/db"
	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":3333")
}

func getEvents(c *gin.Context) {
	events := models.FindAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var e models.Event
	err := c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	e.UserId = 1
	err = e.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, e)
}
