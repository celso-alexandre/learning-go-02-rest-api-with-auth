package routes

import (
	"net/http"
	"strconv"

	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/middlewares"
	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.FindAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var e models.Event
	err := c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwtPayload := middlewares.RetrieveAuthPayload(c)
	e.UserId = jwtPayload.UserId
	err = e.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, e)
}

func getEventById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	e, err := models.FindEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jwtPayload := middlewares.RetrieveAuthPayload(c)
	if e.UserId != jwtPayload.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to access this event"})
		return
	}
	c.JSON(http.StatusOK, e)
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var e models.Event
	err = c.BindJSON(&e)
	jwtPayload := middlewares.RetrieveAuthPayload(c)
	e.UserId = jwtPayload.UserId
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	e.Id = id
	err = e.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, e)
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
