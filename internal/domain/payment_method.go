package domain

// Strategy Interface
type PaymentMethod interface {
	ProcessPayment(amount float64) error
	ValidatePayment() error
	GetMethodName() string
}
