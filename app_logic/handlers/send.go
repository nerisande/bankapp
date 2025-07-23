package handlers

import (
	"bankapp/app_logic/db"
	"bankapp/app_logic/models"
	"bankapp/app_logic/utils"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Обработка запроса /api/send

func Send(w http.ResponseWriter, r *http.Request) {
	// Обрабатываем полученный json
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		utils.ResponseWithError(w, r, err, http.StatusBadRequest)
	}
	// Округляем дробную часть до 2 символов
	transaction.Amount = transaction.Amount.Round(2)
	// Нельзя отправлять деньги самому себе
	if transaction.FromAddress == transaction.ToAddress {
		utils.ResponseWithError(w, r, errors.New("cannot send money to yourself"), http.StatusBadRequest)
		return
	}
	// Проверяем кошелек-отправитель
	var fromWallet models.Wallet
	if err := db.DB.Where("address = ?", transaction.FromAddress).First(&fromWallet).Error; err != nil {
		utils.ResponseWithError(w, r, err, http.StatusBadRequest)
		return
	}
	// Проверяем кошелек-получатель
	var toWallet models.Wallet
	if err := db.DB.Where("address = ?", transaction.ToAddress).First(&toWallet).Error; err != nil {
		utils.ResponseWithError(w, r, err, http.StatusBadRequest)
		return
	}
	// Проверяем баланс - при отрицательном или недостаточном балансе операция отклоняется
	if fromWallet.Balance.Cmp(transaction.Amount) < 0 || fromWallet.Balance.Cmp(decimal.Zero) <= 0 {
		utils.ResponseWithError(w, r, errors.New("incorrect or insufficient balance"), http.StatusBadRequest)
		return
	}
	// Обновляем баланс и сохраняем транзакцию в БД
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		fromWallet.Balance = fromWallet.Balance.Sub(transaction.Amount)
		toWallet.Balance = toWallet.Balance.Add(transaction.Amount)
		if err := tx.Save(&fromWallet).Error; err != nil {
			return err
		}
		if err := tx.Save(&toWallet).Error; err != nil {
			return err
		}
		if err := tx.Save(&transaction).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		utils.ResponseWithError(w, r, err, http.StatusInternalServerError)
		return
	}
	// Отправляем сообщение об успехе и данные о совершенной транзакции
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
