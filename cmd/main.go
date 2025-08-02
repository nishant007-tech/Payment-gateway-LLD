package main

import (
	"fmt"
	"log"
	"payment-gateway/internal/domain"
	"payment-gateway/internal/processor"
)

func main() {
	fmt.Println("🚀 Scalable Payment Gateway - Registry Pattern Demo")
	fmt.Println("==================================================")

	paymentProcessor := processor.NewPaymentProcessor()

	// Show capabilities
	fmt.Printf("📋 Supported Methods: %v\n", paymentProcessor.GetSupportedMethods())
	fmt.Printf("📋 Supported Providers: %v\n", paymentProcessor.GetSupportedProviders())

	// Example 1: UPI Payment
	fmt.Println("\n📱 Example 1: UPI Payment with Razorpay")
	upiDetails := map[string]string{
		"upiID": "user@paytm",
		"pin":   "1234",
	}

	result1, err := paymentProcessor.ProcessPayment(5000, "UPI", "RAZORPAY", upiDetails)
	handleResult(result1, err)

	// Example 2: Card Payment
	fmt.Println("\n💳 Example 2: Card Payment with PayPal")
	cardDetails := map[string]string{
		"cardNumber": "4111111111111111",
		"cvv":        "123",
		"expiryDate": "12/25",
	}
	result2, err := paymentProcessor.ProcessPayment(15000, "CARD", "PAYPAL", cardDetails)
	handleResult(result2, err)

}

func handleResult(result *domain.TransactionResult, err error) {
	if err != nil {
		log.Printf("❌ Payment Error: %v", err)
		if result != nil {
			fmt.Printf("📄 Failed Transaction: %s via %s\n", result.TransactionID, result.Provider)
		}
	} else {
		fmt.Printf("✅ Payment Success!\n")
		fmt.Printf("📄 Transaction ID: %s\n", result.TransactionID)
		fmt.Printf("🌉 Provider: %s\n", result.Provider)
		fmt.Printf("💳 Method: %s\n", result.Method)
		fmt.Printf("💰 Amount: ₹%.2f\n", result.Amount)
	}
}
