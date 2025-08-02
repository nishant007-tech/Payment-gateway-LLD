package processor

import (
	"fmt"
	"payment-gateway/internal/domain"
	"payment-gateway/internal/gateways"
	"payment-gateway/internal/methods"
	"payment-gateway/internal/registry"
)

type PaymentProcessor struct {
	methodRegistry  registry.PaymentMethodRegistry
	gatewayRegistry registry.PaymentGatewayRegistry
}

func NewPaymentProcessor() *PaymentProcessor {
	processor := &PaymentProcessor{
		methodRegistry:  registry.NewPaymentMethodRegistry(),
		gatewayRegistry: registry.NewPaymentGatewayRegistry(),
	}
	// Register all payment methods
	processor.registerPaymentMethods()

	// Register all gateways
	processor.registerGateways()

	return processor
}

// registerPaymentMethods registers all available payment methods
func (pp *PaymentProcessor) registerPaymentMethods() {
	pp.methodRegistry.Register("UPI", methods.CreateUpiPayment)
	pp.methodRegistry.Register("CARD", methods.CreateCardPayment)
}

// registerGateways registers all available payment gateways
func (pp *PaymentProcessor) registerGateways() {
	pp.gatewayRegistry.Register("RAZORPAY", gateways.CreateRazorpayGateway)
	pp.gatewayRegistry.Register("PAYPAL", gateways.CreatePayPalGateway)
}

// ProcessPayment processes a payment using the specified method and provider
func (pp *PaymentProcessor) ProcessPayment(amount float64, method, provider string, details map[string]string) (*domain.TransactionResult, error) {
	fmt.Printf("\n🔄 Processing %s payment via %s...\n", method, provider)
	fmt.Printf("💰 Amount: ₹%.2f\n", amount)

	// Create payment method using registry
	paymentMethod, err := pp.methodRegistry.Create(method, details)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment method: %w", err)
	}

	// Validate payment
	if err := paymentMethod.ValidatePayment(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Create gateway using registry
	gateway, err := pp.gatewayRegistry.Create(provider)
	if err != nil {
		return nil, fmt.Errorf("failed to create gateway: %w", err)
	}

	// Process transaction
	result, err := gateway.ProcessTransaction(paymentMethod, amount)
	if err != nil {
		return result, fmt.Errorf("transaction failed: %w", err)
	}

	fmt.Printf("✅ Transaction completed: %s\n", result.TransactionID)
	return result, nil
}

// RegisterPaymentMethod allows runtime registration of new payment methods
func (pp *PaymentProcessor) RegisterPaymentMethod(methodType string, creator domain.PaymentMethodCreator) {
	pp.methodRegistry.Register(methodType, creator)
}

// RegisterGateway allows runtime registration of new gateways
func (pp *PaymentProcessor) RegisterGateway(provider string, creator domain.PaymentGatewayCreator) {
	pp.gatewayRegistry.Register(provider, creator)
}

// GetSupportedMethods returns all registered payment methods
func (pp *PaymentProcessor) GetSupportedMethods() []string {
	return pp.methodRegistry.GetSupportedMethods()
}

// GetSupportedProviders returns all registered providers
func (pp *PaymentProcessor) GetSupportedProviders() []string {
	return pp.gatewayRegistry.GetSupportedProviders()
}

// IsMethodSupported checks if a payment method is supported
func (pp *PaymentProcessor) IsMethodSupported(method string) bool {
	return pp.methodRegistry.IsSupported(method)
}

// IsProviderSupported checks if a provider is supported
func (pp *PaymentProcessor) IsProviderSupported(provider string) bool {
	return pp.gatewayRegistry.IsSupported(provider)
}
