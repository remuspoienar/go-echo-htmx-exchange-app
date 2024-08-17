package model

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomContext struct {
	Db             *gorm.DB
	CoinbaseApiUrl string
	ExchangeUrl    string
}

type EchoCustomContext struct {
	echo.Context
	CustomContext
}
