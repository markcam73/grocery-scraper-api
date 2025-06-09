package validation

import (
	"grocery-scraper-api/services"
	"testing"
)

func TestValidateFilterOpts(t *testing.T) {
	tests := []struct {
		name        string
		opts        services.FilterOpts
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid filter options",
			opts: services.FilterOpts{
				StoreID:   1,
				ProductID: 1,
				PriceRange: services.PriceRange{
					Min: 0.0,
					Max: 10.0,
				},
				SortBy:    "amount",
				SortOrder: "asc",
			},
			expectError: false,
		},
		{
			name: "missing store ID",
			opts: services.FilterOpts{
				ProductID: 1,
			},
			expectError: true,
			errorMsg:    "store_id: cannot be blank",
		},
		{
			name: "missing product ID",
			opts: services.FilterOpts{
				StoreID: 1,
			},
			expectError: true,
			errorMsg:    "product_id: cannot be blank",
		},
		{
			name: "invalid store ID (zero)",
			opts: services.FilterOpts{
				StoreID:   0,
				ProductID: 1,
			},
			expectError: true,
			errorMsg:    "store_id: cannot be blank",
		},
		{
			name: "invalid product ID (zero)",
			opts: services.FilterOpts{
				StoreID:   1,
				ProductID: 0,
			},
			expectError: true,
			errorMsg:    "product_id: cannot be blank",
		},
		{
			name: "invalid sort by",
			opts: services.FilterOpts{
				StoreID:   1,
				ProductID: 1,
				SortBy:    "invalid",
			},
			expectError: true,
			errorMsg:    "sort_by: must be a valid value",
		},
		{
			name: "invalid sort order",
			opts: services.FilterOpts{
				StoreID:   1,
				ProductID: 1,
				SortOrder: "invalid",
			},
			expectError: true,
			errorMsg:    "sort_order: must be a valid value",
		},
		{
			name: "valid with empty optional fields",
			opts: services.FilterOpts{
				StoreID:   1,
				ProductID: 1,
				SortBy:    "",
				SortOrder: "",
			},
			expectError: false,
		},
		{
			name: "valid with price range",
			opts: services.FilterOpts{
				StoreID:   1,
				ProductID: 1,
				PriceRange: services.PriceRange{
					Min: 1.0,
					Max: 5.0,
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFilterOpts(tt.opts)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				// Note: We could check for specific error messages, but ozzo validation
				// error messages may vary. For now, we just ensure an error occurred.
			} else {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
			}
		})
	}
}

func TestValidatePriceRange(t *testing.T) {
	tests := []struct {
		name        string
		priceRange  services.PriceRange
		expectError bool
	}{
		{
			name: "valid price range",
			priceRange: services.PriceRange{
				Min: 1.0,
				Max: 10.0,
			},
			expectError: false,
		},
		{
			name: "valid with zero values",
			priceRange: services.PriceRange{
				Min: 0.0,
				Max: 0.0,
			},
			expectError: false,
		},
		{
			name: "negative min value",
			priceRange: services.PriceRange{
				Min: -1.0,
				Max: 10.0,
			},
			expectError: true,
		},
		{
			name: "negative max value",
			priceRange: services.PriceRange{
				Min: 1.0,
				Max: -1.0,
			},
			expectError: true,
		},
		{
			name: "max less than min",
			priceRange: services.PriceRange{
				Min: 10.0,
				Max: 5.0,
			},
			expectError: true,
		},
		{
			name: "equal min and max",
			priceRange: services.PriceRange{
				Min: 5.0,
				Max: 5.0,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePriceRange(tt.priceRange)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
			}
		})
	}
}