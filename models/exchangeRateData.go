package models

type ExchangeRateData struct {
	Data []struct {
		CountryCurrencyDesc string `json:"country_currency_desc"`
		ExchangeRate        string `json:"exchange_rate"`
		RecordDate          string `json:"record_date"`
	} `json:"data"`

	Meta struct {
		Count  int `json:"count"`
		Labels struct {
			CountryCurrencyDesc string `json:"country_currency_desc"`
			ExchangeRate        string `json:"exchange_rate"`
			RecordDate          string `json:"record_date"`
		} `json:"labels"`
		DataTypes struct {
			CountryCurrencyDesc string `json:"country_currency_desc"`
			ExchangeRate        string `json:"exchange_rate"`
			RecordDate          string `json:"record_date"`
		} `json:"dataTypes"`
		DataFormats struct {
			CountryCurrencyDesc string `json:"country_currency_desc"`
			ExchangeRate        string `json:"exchange_rate"`
			RecordDate          string `json:"record_date"`
		} `json:"dataFormats"`
		TotalCount int `json:"total-count"`
		TotalPages int `json:"total-pages"`
	} `json:"meta"`

	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}
