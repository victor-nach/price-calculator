package sliceStore

import (
	"github.com/victor-nach/price-calculator/db"
	"github.com/victor-nach/price-calculator/models"
)

type sliceStore struct {
	rates []*models.Rate
}

var _ db.Datastore = &sliceStore{}

func New() *sliceStore {
	store := make([]*models.Rate, 0)
	return &sliceStore{rates: store}
}

func (s *sliceStore)  SaveRate(rate *models.Rate) (*models.Rate, error) {
	s.rates = append(s.rates, rate)
	return rate, nil
}

