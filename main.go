package main

import (
	"github.com/TobiasAagaard/meal-map-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.SetupRoutes(server)
	server.Run(":8080")
}
