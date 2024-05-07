package gateway

import (
	"fmt"

	"github.com/stripe/stripe-go/v78"
)

func HandleStripeErrors(err error) error {
	// Try to safely cast a generic error to a stripe.Error so that we can get at
	// some additional Stripe-specific information about what went wrong.
	if stripeErr, ok := err.(*stripe.Error); ok {
		// The Code field will contain a basic identifier for the failure.
		switch stripeErr.Code {
		case stripe.ErrorCodeCardDeclined:
		case stripe.ErrorCodeExpiredCard:
		case stripe.ErrorCodeIncorrectCVC:
		case stripe.ErrorCodeIncorrectZip:
			// etc.
		}
		// The Err field can be coerced to a more specific error type with a type
		// assertion. This technique can be used to get more specialized
		// information for certain errors.
		if cardErr, ok := stripeErr.Err.(*stripe.CardError); ok {
			return fmt.Errorf("Card was declined with code: %v\n", cardErr.DeclineCode)
		} else {
			return fmt.Errorf("Other Stripe error occurred: %v\n", stripeErr.Error())
		}
	} else {
		return fmt.Errorf("Other error occurred: %v\n", err.Error())
	}
}
