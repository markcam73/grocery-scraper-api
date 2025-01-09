package main

import (
	"go-gin-api/config"
	"go-gin-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize router
	router := gin.Default()

	// Initialize routes
	routes.SetupRoutes(router)

	// Start server
	if err := router.Run(config.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
