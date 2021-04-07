package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/victor-nach/price-calculator/graph/generated"
	"github.com/victor-nach/price-calculator/lib/validator"
	"github.com/victor-nach/price-calculator/models"
)

//CalculatePrice graphql query that returns the price in naira based on a specified exchange rate, tradetyp and margin
func (r *queryResolver) CalculatePrice(ctx context.Context, typeArg models.TradeType, margin float64, exchangeRate float64) (float64, error) {
	err := validator.ValidatePriceInput(margin, exchangeRate)
	if err != nil {
		return 0, err
	}
	currentPrice, err := r.PriceService.GetCurrentPrice()
	if err != nil {
		return 0, err
	}
	cpUSD := currentPrice.Bpi.USD.RateFloat
	var price float64
	switch typeArg {
	case models.TradeTypeBuy:
		price = cpUSD - margin*cpUSD
	case models.TradeTypeSell:
		price = cpUSD - margin*cpUSD
	}

	fmt.Println(cpUSD, price)
	// get price in Naira
	price *= exchangeRate
	fmt.Println(cpUSD, price)
	return price, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
