package service

import (
	"goro/db"
	"goro/model"
	"os"
)

func GetContext() model.CustomContext {
	d := db.New()
	coinbaseApiUrl := os.Getenv("COINBASE_CURRENCIES_URL")
	exchangeUrl := os.Getenv("COINBASE_EXCHANGE_URL")

	return model.CustomContext{
		Db:             d,
		CoinbaseApiUrl: coinbaseApiUrl,
		ExchangeUrl:    exchangeUrl,
	}
}
