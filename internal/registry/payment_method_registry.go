package registry

import (
	"errors"
	"payment-gateway/internal/domain"
	"strings"
)

// PaymentMethodRegistry manages payment method creators using Registry pattern

type PaymentMethodRegistry struct {
	creators map[string]domain.PaymentMethodCreator // e.g upi, paypal
}

func NewPaymentMethodRegistry() PaymentMethodRegistry {
	return PaymentMethodRegistry{
		creators: make(map[string]domain.PaymentMethodCreator),
	}
}

func (pmr *PaymentMethodRegistry) Register(methodType string, creator domain.PaymentMethodCreator) {
	pmr.creators[strings.ToUpper(methodType)] = creator
}

func (pmr *PaymentMethodRegistry) Create(methodType string, details map[string]string) (domain.PaymentMethod, error) {
	creator, exists := pmr.creators[strings.ToUpper(methodType)]
	if !exists {
		return nil, errors.New("unsupported payment method: " + methodType)
	}
	return creator(details)
}

// GetSupportedMethods returns all registered payment method types
func (pmr *PaymentMethodRegistry) GetSupportedMethods() []string {
	methods := make([]string, 0, len(pmr.creators))
	for method := range pmr.creators {
		methods = append(methods, method)
	}
	return methods
}

// Unregister removes a payment method from the registry
func (pmr *PaymentMethodRegistry) Unregister(methodType string) {
	delete(pmr.creators, strings.ToUpper(methodType))
}

// IsSupported checks if a payment method is registered
func (pmr *PaymentMethodRegistry) IsSupported(methodType string) bool {
	_, exists := pmr.creators[strings.ToUpper(methodType)]
	return exists
}
