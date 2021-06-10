package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/victor-nach/price-calculator/graph/generated"
	"github.com/victor-nach/price-calculator/lib/validator"
	"github.com/victor-nach/price-calculator/models"
)

func (r *queryResolver) CalculatePrice(ctx context.Context, typeArg models.TradeType, margin float64, exchangeRate float64) (float64, error) {
	err := validator.ValidatePriceInput(margin, exchangeRate)
	if err != nil {
		return 0, err
	}
	price, err := r.calcPrice(typeArg, margin, exchangeRate)
	if err != nil {
		return 0, nil
	}
	return price, nil
}

func (r *Resolver) calcPrice(typeArg models.TradeType, margin float64, exchangeRate float64) (float64, error) {
	currentPrice, err := r.PriceService.GetCurrentPrice()
	if err != nil {
		return 0, err
	}

	cpUSD := currentPrice.Bpi.USD.RateFloat
	var price float64

	switch typeArg {
	case models.TradeTypeBuy:
		price = cpUSD - (margin/100)*cpUSD
	case models.TradeTypeSell:
		price = cpUSD - (margin/100)*cpUSD
	}

	// get price in Naira
	price *= exchangeRate
	rate := models.Rate{
		ID:        r.idGen.Generate(),
		Exchange:  exchangeRate,
		Price:     price,
		TradeType: typeArg.String(),
		Ts:        time.Now(),
	}
	_, err = r.Store.SaveRate(&rate)
	if err != nil {
		return 0, err
	}
	return price, nil
}

func (r *subscriptionResolver) CalculatePrice(ctx context.Context, typeArg models.TradeType, margin float64, exchangeRate float64) (<-chan float64, error) {
	err := validator.ValidatePriceInput(margin, exchangeRate)
	if err != nil {
		return nil, err
	}
	respChan := make(chan float64, 1)
	go func() {
		for {
			price, err := r.calcPrice(typeArg, margin, exchangeRate)
			if err != nil {
				log.Println(err)
				break
			}
			fmt.Println(price)
			respChan <- -price + 1
			time.Sleep(10 * time.Second)
		}
	}()
	return respChan, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
