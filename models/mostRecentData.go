package models

import "time"

type MostRecentData struct {
	CountryCurrencyDesc string
	ExchangeRate        float64
	RecordDate          time.Time
}
