package validation

import (
	"grocery-scraper-api/services"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ValidateFilterOpts validates the FilterOpts struct using ozzo validation
func ValidateFilterOpts(opts services.FilterOpts) error {
	return validation.ValidateStruct(&opts,
		validation.Field(&opts.StoreID, validation.Required, validation.Min(uint32(1))),
		validation.Field(&opts.ProductID, validation.Required, validation.Min(uint32(1))),
		validation.Field(&opts.PriceRange, validation.By(validatePriceRange)),
		validation.Field(&opts.SortBy, validation.In("amount", "store", "item", "")),
		validation.Field(&opts.SortOrder, validation.In("asc", "desc", "")),
	)
}

// ValidatePriceRange validates the PriceRange struct
func ValidatePriceRange(opts services.PriceRange) error {
	return validation.ValidateStruct(&opts,
		validation.Field(&opts.Min, validation.Min(0.0)),
		validation.Field(&opts.Max, validation.Min(0.0)),
		validation.Field(&opts.Max, validation.When(opts.Min > 0 && opts.Max > 0, validation.Min(opts.Min))),
	)
}

// validatePriceRange is a custom validation function for PriceRange
func validatePriceRange(value interface{}) error {
	priceRange, ok := value.(services.PriceRange)
	if !ok {
		return validation.NewError("validation_price_range", "must be a valid price range")
	}

	return ValidatePriceRange(priceRange)
}
