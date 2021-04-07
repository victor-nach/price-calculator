package graph

import "github.com/victor-nach/price-calculator/lib"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver defines all the dependencies required by the resolver handlers
type Resolver struct {
	PriceService lib.PriceService
}

// NewResolver returns a new resolver
func NewResolver(priceService lib.PriceService) *Resolver {
	return &Resolver{PriceService: priceService}
}
