package gateways

import (
	"fmt"
	"payment-gateway/internal/domain"
	"payment-gateway/internal/utils"
)

type PayPalGateway struct {
	ClientID     string
	ClientSecret string
}

func CreatePayPalGateway() domain.PaymentGateway {
	return &PayPalGateway{
		ClientID:     "paypal_client_id",
		ClientSecret: "paypal_client_secret",
	}
}

func (p *PayPalGateway) ProcessTransaction(method domain.PaymentMethod, amount float64) (*domain.TransactionResult, error) {
	fmt.Println("🌉 Processing via PayPal Gateway...")

	err := method.ProcessPayment(amount)
	if err != nil {
		return &domain.TransactionResult{
			Success:  false,
			Message:  err.Error(),
			Amount:   amount,
			Provider: "PAYPAL",
			Method:   method.GetMethodName(),
		}, err
	}

	return &domain.TransactionResult{
		Success:       true,
		TransactionID: utils.GenerateTransactionID("PP"),
		Message:       "Payment successful via PayPal",
		Amount:        amount,
		Provider:      "PAYPAL",
		Method:        method.GetMethodName(),
	}, nil
}

func (p *PayPalGateway) GetProviderName() string {
	return "PAYPAL"
}
