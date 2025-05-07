package main

func main() {
	server := gin.Default()
	routes.SetupRoutes(server)
	server.Run(":8080")
}
