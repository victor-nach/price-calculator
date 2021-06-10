package db

import "github.com/victor-nach/price-calculator/models"

//Datastore defines the required store methods
type Datastore interface {
	SaveRate(rate *models.Rate) (*models.Rate, error)
}
