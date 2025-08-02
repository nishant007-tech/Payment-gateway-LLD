package methods

import (
	"errors"
	"fmt"
	"payment-gateway/internal/domain"
)

type UpiPayment struct {
	UpiID string
	Pin   string
}

func CreateUpiPayment(details map[string]string) (domain.PaymentMethod, error) {
	upiID := details["upiID"]
	pin := details["pin"]
	if upiID == "" || pin == "" {
		return nil, errors.New("UPI ID and PIN are required")
	}
	return &UpiPayment{UpiID: upiID, Pin: pin}, nil
}

func (upi *UpiPayment) GetMethodName() string {
	return "UPI"
}

func (upi *UpiPayment) ProcessPayment(amount float64) error {
	fmt.Printf("Processing UPI payment of ₹%.2f via %s\n", amount, upi.UpiID)
	if amount > 100000 {
		return errors.New("UPI limit exceeded (₹1,00,000)")
	}
	return nil
}

func (upi *UpiPayment) ValidatePayment() error {
	if upi.UpiID == "" || upi.Pin == "" {
		return errors.New("UPI details required")
	}
	return nil
}
