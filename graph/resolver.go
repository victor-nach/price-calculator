package graph

import (
	"github.com/victor-nach/price-calculator/db"
	"github.com/victor-nach/price-calculator/lib"
	"github.com/victor-nach/price-calculator/lib/ulid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver defines all the dependencies required by the resolver handlers
type Resolver struct {
	PriceService lib.PriceService
	Store        db.Datastore
	idGen        ulid.Idgenerator
}

// NewResolver returns a new resolver
func NewResolver(priceService lib.PriceService, store db.Datastore) *Resolver {
	return &Resolver{
		PriceService: priceService,
		Store:        store,
		idGen:        ulid.New(),
	}
}
