package db

// QuoteID -> quoteData
var userWallet map[string]map[string]float64 // UserID -> wallet

func InitWalletDB() {
	userWallet = make(map[string]map[string]float64)
}

func TopUpWallet(accountID string, currency string, amount float64) {
	wallet, isExist := GetWallet(accountID)
	if !isExist { // If account doesn't exist, create new account
		wallet = map[string]float64{
			currency: amount,
		}
	} else {
		wallet[currency] += amount
	}
	UpdateWallet(accountID, wallet)
}

func GetWallet(accountID string) (map[string]float64, bool) {
	wallet, isExist := userWallet[accountID]
	return wallet, isExist
}

func UpdateWallet(accountID string, wallet map[string]float64) {
	userWallet[accountID] = wallet
}
