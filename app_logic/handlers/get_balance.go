package handlers

import (
	"bankapp/app_logic/db"
	"bankapp/app_logic/models"
	"bankapp/app_logic/utils"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// Обработка запроса GET /api/wallet/{address}/balance

func GetBalance(w http.ResponseWriter, r *http.Request) {
	// Получаем адрес из /api/wallet/{address}/balance
	address := mux.Vars(r)["address"]
	if address == "" {
		utils.ResponseWithError(w, r, errors.New("missing wallet address"), http.StatusBadRequest)
		return
	}
	// Ищем кошелек в БД
	var wallet models.Wallet
	if err := db.DB.Where("address = ?", address).First(&wallet).Error; err != nil {
		utils.ResponseWithError(w, r, errors.New("invalid wallet address"), http.StatusNotFound)
		return
	}
	// Отправляем кошелек в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallet)
}
