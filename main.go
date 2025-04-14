package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/api/v1")
	{
		v1.GET("/meals")
	}
	server.Run(":8080")
}
