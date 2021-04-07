package validator

import (
	"github.com/victor-nach/price-calculator/lib/rerrors"
	"math"
)

//ValidatePriceInput ....
func ValidatePriceInput(margin float64, exchangeRate float64) error {
	// check if the margin is a valid percentage
	if math.Signbit(margin) || margin > 100 {
		return rerrors.LogFormat(rerrors.InvalidRequestErr, nil)
	}

	// check if the exchangerate is valid
	if math.Signbit(exchangeRate) {
		return rerrors.LogFormat(rerrors.InvalidRequestErr, nil)
	}

	return nil
}
