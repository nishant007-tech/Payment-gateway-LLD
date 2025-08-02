package domain

type TransactionResult struct {
	Success       bool    `json:"success"`
	TransactionID string  `json:"transaction_id"`
	Message       string  `json:"message"`
	Amount        float64 `json:"amount"`
	Provider      string  `json:"provider"`
	Method        string  `json:"method"`
}
