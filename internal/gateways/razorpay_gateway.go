package gateways

import (
	"fmt"
	"payment-gateway/internal/domain"
	"payment-gateway/internal/utils"
)

type RazorpayGateway struct {
	APIKey    string
	APISecret string
}

func CreateRazorpayGateway() domain.PaymentGateway {
	return &RazorpayGateway{
		APIKey:    "rzp_test_key",
		APISecret: "rzp_test_secret",
	}
}

func (rp *RazorpayGateway) ProcessTransaction(method domain.PaymentMethod, amount float64) (*domain.TransactionResult, error) {
	fmt.Println("🌉 Processing via Razorpay Gateway...")

	err := method.ProcessPayment(amount)
	if err != nil {
		return &domain.TransactionResult{
			Success:  false,
			Message:  err.Error(),
			Amount:   amount,
			Provider: "RAZORPAY",
			Method:   method.GetMethodName(),
		}, err
	}

	return &domain.TransactionResult{
		Success:       true,
		TransactionID: utils.GenerateTransactionID("RZP"),
		Message:       "Payment successful via Razorpay",
		Amount:        amount,
		Provider:      "RAZORPAY",
		Method:        method.GetMethodName(),
	}, nil
}

func (r *RazorpayGateway) GetProviderName() string {
	return "RAZORPAY"
}
