package controllers

import (
	"encoding/json"
	"grocery-scraper-api/models"
	"grocery-scraper-api/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetPrices(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedData   []models.Price
		expectError    bool
	}{
		{
			name:           "valid request with required params",
			queryParams:    "store_id=1&product_id=1",
			expectedStatus: http.StatusOK,
			expectedData: []models.Price{
				{ID: 1, Item: "Apple", Store: "Store A", Amount: 0.99, Unit: "kg"},
				{ID: 2, Item: "Banana", Store: "Store B", Amount: 0.59, Unit: "kg"},
				{ID: 3, Item: "Orange", Store: "Store A", Amount: 1.29, Unit: "kg"},
				{ID: 4, Item: "Milk", Store: "Store C", Amount: 1.49, Unit: "L"},
				{ID: 5, Item: "Bread", Store: "Store B", Amount: 2.49, Unit: "loaf"},
			},
			expectError: false,
		},
		{
			name:           "valid request with optional params",
			queryParams:    "store_id=1&product_id=1&sort_by=amount&sort_order=asc",
			expectedStatus: http.StatusOK,
			expectedData: []models.Price{
				{ID: 1, Item: "Apple", Store: "Store A", Amount: 0.99, Unit: "kg"},
				{ID: 2, Item: "Banana", Store: "Store B", Amount: 0.59, Unit: "kg"},
				{ID: 3, Item: "Orange", Store: "Store A", Amount: 1.29, Unit: "kg"},
				{ID: 4, Item: "Milk", Store: "Store C", Amount: 1.49, Unit: "L"},
				{ID: 5, Item: "Bread", Store: "Store B", Amount: 2.49, Unit: "loaf"},
			},
			expectError: false,
		},
		{
			name:           "missing store_id",
			queryParams:    "product_id=1",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "missing product_id",
			queryParams:    "store_id=1",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "invalid store_id",
			queryParams:    "store_id=0&product_id=1",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "invalid product_id",
			queryParams:    "store_id=1&product_id=0",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "invalid sort_by",
			queryParams:    "store_id=1&product_id=1&sort_by=invalid",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "invalid sort_order",
			queryParams:    "store_id=1&product_id=1&sort_order=invalid",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new controller for each test
			controller := NewPriceController()

			// Create a new Gin router
			router := gin.New()
			router.GET("/prices", controller.GetPrices)

			// Create request
			req, err := http.NewRequest(http.MethodGet, "/prices?"+tt.queryParams, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(rr, req)

			// Check status code
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			// Parse response
			var response map[string]interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("could not parse response: %v", err)
			}

			if tt.expectError {
				// Check that error field exists
				if _, exists := response["error"]; !exists {
					t.Errorf("expected error field in response")
				}
			} else {
				// Check success response structure
				if message, exists := response["message"]; !exists || message != "Get all prices" {
					t.Errorf("expected message 'Get all prices', got %v", message)
				}

				// Check data field exists
				data, exists := response["data"]
				if !exists {
					t.Errorf("expected data field in response")
					return
				}

				// Convert data to slice of prices for comparison
				dataBytes, err := json.Marshal(data)
				if err != nil {
					t.Fatalf("could not marshal data: %v", err)
				}

				var actualPrices []models.Price
				err = json.Unmarshal(dataBytes, &actualPrices)
				if err != nil {
					t.Fatalf("could not unmarshal prices: %v", err)
				}

				// Check that we got the expected number of prices
				if len(actualPrices) != len(tt.expectedData) {
					t.Errorf("expected %d prices, got %d", len(tt.expectedData), len(actualPrices))
					return
				}

				// Check each price matches expected data
				for i, expectedPrice := range tt.expectedData {
					if i >= len(actualPrices) {
						t.Errorf("missing price at index %d", i)
						continue
					}

					actualPrice := actualPrices[i]
					if actualPrice.ID != expectedPrice.ID {
						t.Errorf("price %d: expected ID %d, got %d", i, expectedPrice.ID, actualPrice.ID)
					}
					if actualPrice.Item != expectedPrice.Item {
						t.Errorf("price %d: expected Item %s, got %s", i, expectedPrice.Item, actualPrice.Item)
					}
					if actualPrice.Store != expectedPrice.Store {
						t.Errorf("price %d: expected Store %s, got %s", i, expectedPrice.Store, actualPrice.Store)
					}
					if actualPrice.Amount != expectedPrice.Amount {
						t.Errorf("price %d: expected Amount %f, got %f", i, expectedPrice.Amount, actualPrice.Amount)
					}
					if actualPrice.Unit != expectedPrice.Unit {
						t.Errorf("price %d: expected Unit %s, got %s", i, expectedPrice.Unit, actualPrice.Unit)
					}
				}
			}
		})
	}
}

func TestNewPriceController(t *testing.T) {
	controller := NewPriceController()

	if controller == nil {
		t.Error("expected non-nil controller")
		return
	}

	if controller.PriceService == nil {
		t.Error("expected non-nil PriceService")
	}
}

func TestPriceControllerIntegration(t *testing.T) {
	// Test the integration between controller and service
	controller := NewPriceController()

	// Test that the service returns the expected dummy data
	filterOpts := services.FilterOpts{
		StoreID:   1,
		ProductID: 1,
	}

	prices := controller.PriceService.LoadPriceData(filterOpts)

	expectedPrices := []models.Price{
		{ID: 1, Item: "Apple", Store: "Store A", Amount: 0.99, Unit: "kg"},
		{ID: 2, Item: "Banana", Store: "Store B", Amount: 0.59, Unit: "kg"},
		{ID: 3, Item: "Orange", Store: "Store A", Amount: 1.29, Unit: "kg"},
		{ID: 4, Item: "Milk", Store: "Store C", Amount: 1.49, Unit: "L"},
		{ID: 5, Item: "Bread", Store: "Store B", Amount: 2.49, Unit: "loaf"},
	}

	if len(prices) != len(expectedPrices) {
		t.Errorf("expected %d prices, got %d", len(expectedPrices), len(prices))
		return
	}

	for i, expectedPrice := range expectedPrices {
		actualPrice := prices[i]
		if actualPrice.ID != expectedPrice.ID {
			t.Errorf("price %d: expected ID %d, got %d", i, expectedPrice.ID, actualPrice.ID)
		}
		if actualPrice.Item != expectedPrice.Item {
			t.Errorf("price %d: expected Item %s, got %s", i, expectedPrice.Item, actualPrice.Item)
		}
		if actualPrice.Store != expectedPrice.Store {
			t.Errorf("price %d: expected Store %s, got %s", i, expectedPrice.Store, actualPrice.Store)
		}
		if actualPrice.Amount != expectedPrice.Amount {
			t.Errorf("price %d: expected Amount %f, got %f", i, expectedPrice.Amount, actualPrice.Amount)
		}
		if actualPrice.Unit != expectedPrice.Unit {
			t.Errorf("price %d: expected Unit %s, got %s", i, expectedPrice.Unit, actualPrice.Unit)
		}
	}
}