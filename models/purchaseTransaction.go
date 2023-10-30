package models

import (
	"math"
	"time"
)

type PurchaseTransaction struct {
	ID              int64     `json:"id"`
	Description     string    `json:"description" validate:"max=50"`
	PurchaseAmount  float64   `json:"purchase_amount" validate:"min=0.01"`
	TransactionDate time.Time `json:"transaction_date"`
}

func (transaction *PurchaseTransaction) RoundToNearestCent() {
	transaction.PurchaseAmount = math.Round(transaction.PurchaseAmount*100) / 100
}
