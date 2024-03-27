package routes

import (
	"net/http"
	"strconv"

	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/middlewares"
	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwtPayload := middlewares.RetrieveAuthPayload(c)
	register := models.Register{
		EventId: eventId,
		UserId:  jwtPayload.UserId,
	}
	err = register.CreateRegister()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, register)
}

func unregisterForEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwtPayload := middlewares.RetrieveAuthPayload(c)
	register := models.Register{
		EventId: eventId,
		UserId:  jwtPayload.UserId,
	}
	err = register.DeleteRegister()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, register)
}
