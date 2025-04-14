package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {
	v1 := server.Group("/api/v1")
	{
		v1.GET("/meals")
		v1.POST("/meals")

	}
}
