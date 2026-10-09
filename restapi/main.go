package main

import (
	routes "example.com/restapi/api-test/routes"
	"example.com/restapi/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
