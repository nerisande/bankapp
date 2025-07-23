package handlers

import (
	"bankapp/app_logic/db"
	"bankapp/app_logic/models"
	"bankapp/app_logic/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Обработка запроса GET /api/transactions?count=N

func GetLast(w http.ResponseWriter, r *http.Request) {
	// Обработка URL-параметра
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil || count <= 0 {
		utils.ResponseWithError(w, r, errors.New("invalid or missing count parameter"), http.StatusBadRequest)
		return
	}
	var transactions []models.Transaction
	// Получаем последние N транзакций
	if err := db.DB.Order("ID desc").Limit(count).Find(&transactions).Error; err != nil {
		utils.ResponseWithError(w, r, err, http.StatusInternalServerError)
		return
	}
	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
