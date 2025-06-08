package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PriceController struct {
	// TODO: Add price service when implemented
}

func NewPriceController() *PriceController {
	return &PriceController{}
}

func (pc *PriceController) GetPrices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all prices"})
}

func (pc *PriceController) CreatePrice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create new price"})
}

func (pc *PriceController) GetPricesByStore(c *gin.Context) {
	storeID := c.Param("store_id")
	c.JSON(http.StatusOK, gin.H{"message": "Get prices by store", "store_id": storeID})
}

func (pc *PriceController) GetPricesByProduct(c *gin.Context) {
	productID := c.Param("product_id")
	c.JSON(http.StatusOK, gin.H{"message": "Get prices by product", "product_id": productID})
}
