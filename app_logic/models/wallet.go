package models

import "github.com/shopspring/decimal"

// Структура для описания кошелька
type Wallet struct {
	ID uint `gorm:"index"`
	// Адрес кошелька
	Address string `gorm:"unique" json:"address"`
	// Баланс; Используем формат decimal вместо float дабы избежать возможных ошибок с округлением
	Balance decimal.Decimal `json:"balance"`
}
