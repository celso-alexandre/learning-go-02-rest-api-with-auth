package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/ping", getEvents)

	server.Run(":3333")
}

func getEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
