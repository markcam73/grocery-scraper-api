package services

import (
	"grocery-scraper-api/models"
)

type PriceService struct {
	prices []models.Price
}

// FilterOpts defines the options for filtering and sorting prices.
type FilterOpts struct {
	StoreID    uint32     `form:"store_id" json:"store_id"`
	ProductID  uint32     `form:"product_id" json:"product_id"`
	PriceRange PriceRange `form:"price_range" json:"price_range"`
	SortBy     string     `form:"sort_by" json:"sort_by"`       // e.g., "amount", "store"
	SortOrder  string     `form:"sort_order" json:"sort_order"` // e.g., "asc", "desc"
}

type PriceRange struct {
	Min float64
	Max float64
}

// PriceService provides methods to manage and retrieve price data.

func NewPriceService() *PriceService {
	return &PriceService{
		prices: make([]models.Price, 0),
	}
}

func (s *PriceService) GetPrices() []models.Price {
	return s.prices
}

func (s *PriceService) LoadPriceData(filterOpts FilterOpts) []models.Price {
	// This method would typically load price data from a database or an external API.
	// For now, we'll simulate loading data with some dummy prices.
	s.prices = []models.Price{
		{ID: 1, Item: "Apple", Store: "Store A", Amount: 0.99, Unit: "kg"},
		{ID: 2, Item: "Banana", Store: "Store B", Amount: 0.59, Unit: "kg"},
		{ID: 3, Item: "Orange", Store: "Store A", Amount: 1.29, Unit: "kg"},
		{ID: 4, Item: "Milk", Store: "Store C", Amount: 1.49, Unit: "L"},
		{ID: 5, Item: "Bread", Store: "Store B", Amount: 2.49, Unit: "loaf"},
	}
	return s.prices
}
