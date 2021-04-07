package coindesk

import (
	"encoding/json"
	"github.com/victor-nach/price-calculator/lib/rerrors"
	"github.com/victor-nach/price-calculator/models"
	"io/ioutil"
	"net/http"
	"time"
)

type coindesk struct {
	baseURL    string
	httpClient *http.Client
}

// ensure coindesk client implements price service interface

// NewClient returns a new coindesk client
func NewClient(baseUrl string) *coindesk {
	return &coindesk{
		baseURL: baseUrl,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetCurrentPrice returns the current price of bitcoin
func (c *coindesk) GetCurrentPrice() (*models.CurrentPrice, error) {
	resp, err := c.httpClient.Get(c.baseURL)
	if err != nil {
		return nil, rerrors.LogFormat(rerrors.InternalErr, err)
	}

	// use keep-alive method for persistent connections
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, rerrors.LogFormat(rerrors.InternalErr, err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, rerrors.LogFormat(rerrors.InternalErr, err)
	}

	var currentPrice models.CurrentPrice
	err = json.Unmarshal(body, &currentPrice)
	if err != nil {
		return nil, rerrors.LogFormat(rerrors.InternalErr, err)
	}
	return &currentPrice, nil
}
