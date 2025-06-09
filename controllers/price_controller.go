package controllers

import (
	"grocery-scraper-api/services"
	"grocery-scraper-api/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PriceController struct {
	PriceService *services.PriceService
}

func NewPriceController() *PriceController {
	return &PriceController{
		PriceService: services.NewPriceService(),
	}
}

func (pc *PriceController) GetPrices(c *gin.Context) {
	// load params
	var filterOpts services.FilterOpts
	if err := c.ShouldBindQuery(&filterOpts); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// validate params
	if err := validation.ValidateFilterOpts(filterOpts); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	prices := pc.PriceService.LoadPriceData(filterOpts)
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all prices",
		"data":    prices,
	})
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
