package routes

import (
	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/events", getEvents)

	// Authenticated routes
	// server.Use(middlewares.AuthMiddleware())
	authenticated := server.Group("/")
	authenticated.Use(middlewares.AuthMiddleware())

	authenticated.GET("/events/:id", getEventById)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
}
