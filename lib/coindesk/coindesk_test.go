package coindesk

import (
	"github.com/stretchr/testify/assert"
	"github.com/victor-nach/price-calculator/config"
	"testing"
)

func TestCoindesk_GetCurrentPrice(t *testing.T) {
	// use default env
	cfg := config.LoadSecrets()
	coindeskClient := NewClient(cfg.CoindeskURL)
	cp, err := coindeskClient.GetCurrentPrice()
	assert.NoError(t, err)
	assert.NotNil(t, cp)
	assert.NotNil(t, cp.Bpi.EUR)
	assert.NotNil(t, cp.Bpi.GBP)
	assert.NotNil(t, cp.Bpi.USD)
	assert.NotNil(t, cp.Bpi.USD.Rate)
}
