package models

import "time"

type PurchaseTransactionOutput struct {
	ID              int64     `json:"id"`
	Description     string    `json:"description" validate:"max=50"`
	PurchaseAmount  float64   `json:"purchase_amount" validate:"min=0.01"`
	TransactionDate time.Time `json:"transaction_date"`
	ExchangeRate    float64   `json:"exchange_rate"`
	ConvertedAmount float64   `json:"converted_amount"`
}
