package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {
	api := server.Group("/api/v1")
	{
		api.GET("/meals")

	}
}
