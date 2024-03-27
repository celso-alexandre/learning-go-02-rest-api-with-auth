package routes

import (
	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	// Protected routes
	server.Use(middlewares.AuthMiddleware())

	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
