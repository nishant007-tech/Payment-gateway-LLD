package domain

type PaymentGateway interface {
	ProcessTransaction(method PaymentMethod, amount float64) (*TransactionResult, error)
	GetProviderName() string
}
