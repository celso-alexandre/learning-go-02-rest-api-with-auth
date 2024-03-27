package main

import (
	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/db"
	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":3333")
}
