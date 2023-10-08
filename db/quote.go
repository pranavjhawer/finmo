package db

import "time"

type QuoteData struct {
	ConversionRate float64
	FromCurrency   string
	ToCurrency     string
	ExpiryAt       time.Time
}

var quoteIDRate map[string]QuoteData

func InitQuoteDB() {
	quoteIDRate = make(map[string]QuoteData)
}

func SetQuote(quoteID string, quote QuoteData) {
	quoteIDRate[quoteID] = quote
}

func GetQuote(quoteID string) (QuoteData, bool) {
	quote, isExist := quoteIDRate[quoteID]
	return quote, isExist
}
