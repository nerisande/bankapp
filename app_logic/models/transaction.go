package models

import "github.com/shopspring/decimal"

// Структура для описания транзакции
type Transaction struct {
	ID uint `gorm:"index"`
	// Адрес отправителя
	FromAddress string `json:"from"`
	// Адрес получателя
	ToAddress string `json:"to"`
	// Сумма перевода; Используем формат decimal вместо float дабы избежать возможных ошибок с округлением
	Amount decimal.Decimal `json:"amount"`
}
