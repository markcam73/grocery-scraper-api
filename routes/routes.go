package routes

import (
	"grocery-scraper-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Create controllers
	userController := controllers.NewUserController()
	priceController := controllers.NewPriceController()

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		{
			users.GET("/", userController.GetUsers)
			users.POST("/", userController.CreateUser)
		}

		// Price routes
		prices := v1.Group("/prices")
		{
			prices.GET("/", priceController.GetPrices)
			prices.POST("/", priceController.CreatePrice)
			prices.GET("/store/:store_id", priceController.GetPricesByStore)
			prices.GET("/product/:product_id", priceController.GetPricesByProduct)
		}
	}
}
