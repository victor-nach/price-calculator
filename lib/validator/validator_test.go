package validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/victor-nach/price-calculator/lib/rerrors"
	"testing"
)

func TestValidatePriceInput(t *testing.T) {
	type tests struct {
		margin       float64
		exchangeRate float64
		err          error
	}
	cases := []struct {
		margin       float64
		exchangeRate float64
		err          error
	}{
		{margin: 0.1, exchangeRate: 598, err: nil},
		{margin: -0.6, exchangeRate: 38, err: rerrors.LogFormat(rerrors.InvalidRequestErr, nil)},
		{margin: 0.143, exchangeRate: -478, err: rerrors.LogFormat(rerrors.InvalidRequestErr, nil)},
	}
	for idx, next := range cases {
		t.Run(fmt.Sprintf("subtest-%d", idx), func(t *testing.T) {
			err := ValidatePriceInput(next.margin, next.exchangeRate)
			assert.Equal(t, err, next.err)
		})
	}
}
