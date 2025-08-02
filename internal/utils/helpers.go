package utils

import (
	"fmt"
	"time"
)

// GenerateTransactionID generates a unique transaction ID with prefix
func GenerateTransactionID(prefix string) string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%s_%d_%d", prefix, timestamp, generateRandomNumber())
}

// generateRandomNumber generates a simple random number
func generateRandomNumber() int64 {
	return time.Now().UnixNano() % 1000000
}

type CurrencyFormats interface {
	CurrencyFormat() string
}

type INRCurrency struct{}

func (inr *INRCurrency) CurrencyFormat() string {
	return "₹"
}

type USDCurrency struct{}

func (usd *USDCurrency) CurrencyFormat() string {
	return "$"
}

// FormatAmount formats amount for display
func FormatAmount(amount float64, currency string) string {
	switch currency {
	case "INR":
		return fmt.Sprintf("₹%.2f", amount)
	case "USD":
		return fmt.Sprintf("$%.2f", amount)
	default:
		return fmt.Sprintf("%.2f %s", amount, currency)
	}
}
