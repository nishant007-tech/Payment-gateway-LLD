package methods

import (
	"errors"
	"fmt"
	"payment-gateway/internal/domain"
)

type CardPayment struct {
	CardNumber string
	CVV        string
	ExpiryDate string
}

func CreateCardPayment(details map[string]string) (domain.PaymentMethod, error) {
	if details["cardNumber"] == "" || details["cvv"] == "" || details["expiryDate"] == "" {
		return nil, errors.New("card number, CVV, and expiry date are required")
	}
	return &CardPayment{
		CardNumber: details["cardNumber"],
		CVV:        details["cvv"],
		ExpiryDate: details["expiryDate"],
	}, nil
}

func (c *CardPayment) ProcessPayment(amount float64) error {
	fmt.Printf("Processing Card payment of ₹%.2f via %s\n", amount, c.CardNumber)
	if amount > 500000 {
		return errors.New("card limit exceeded (₹5,00,000)")
	}
	return nil
}

func (c *CardPayment) ValidatePayment() error {
	if len(c.CardNumber) < 16 || c.CVV == "" {
		return errors.New("invalid card details")
	}
	return nil
}

func (c *CardPayment) GetMethodName() string {
	return "CARD"
}
