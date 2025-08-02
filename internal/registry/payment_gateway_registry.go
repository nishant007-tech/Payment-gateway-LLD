package registry

import (
	"errors"
	"payment-gateway/internal/domain"
	"strings"
)

// PaymentGatewayRegistry manages payment gateway creators using Registry pattern
type PaymentGatewayRegistry struct {
	creators map[string]domain.PaymentGatewayCreator
}

func NewPaymentGatewayRegistry() PaymentGatewayRegistry {
	return PaymentGatewayRegistry{
		creators: make(map[string]domain.PaymentGatewayCreator),
	}
}

func (pgr *PaymentGatewayRegistry) Register(provider string, creator domain.PaymentGatewayCreator) {
	pgr.creators[strings.ToUpper(provider)] = creator
}

func (pgr *PaymentGatewayRegistry) Create(provider string) (domain.PaymentGateway, error) {
	creator, exists := pgr.creators[strings.ToUpper(provider)]
	if !exists {
		return nil, errors.New("unsupported provider: " + provider)
	}
	return creator(), nil
}

// GetSupportedProviders returns all registered provider names
func (pgr *PaymentGatewayRegistry) GetSupportedProviders() []string {
	providers := make([]string, 0, len(pgr.creators))
	for provider := range pgr.creators {
		providers = append(providers, provider)
	}
	return providers
}

// Unregister removes a payment gateway from the registry
func (pgr *PaymentGatewayRegistry) Unregister(provider string) {
	delete(pgr.creators, strings.ToUpper(provider))
}

// IsSupported checks if a payment gateway is registered
func (pgr *PaymentGatewayRegistry) IsSupported(provider string) bool {
	_, exists := pgr.creators[strings.ToUpper(provider)]
	return exists
}
