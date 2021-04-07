package lib

import "github.com/victor-nach/price-calculator/models"

type PriceService interface {
	GetCurrentPrice() (*models.CurrentPrice, error)
}
