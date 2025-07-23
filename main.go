package main

import (
	"bankapp/app_logic/db"
	"bankapp/app_logic/handlers"
	"bankapp/app_logic/models"
	"bankapp/app_logic/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)

func main() {
	// Инициализируем БД
	DB, err := db.Init_db()
	if err != nil {
		log.Fatal(DB, err)
	}
	// Если БД пустая - создаём 10 кошельков и выводим их в консоль для упрощения тестирования
	var count int64
	if err := DB.Model(&models.Wallet{}).Count(&count).Error; err != nil {
		log.Fatal(err)
		return
	}
	if count == 0 {
		for i := 0; i < 10; i++ {
			newWallet := models.Wallet{
				Address: utils.GenerateRandomAddress(),
				Balance: decimal.NewFromFloat(100.0),
			}
			result := DB.Create(&newWallet)
			if result.Error != nil {
				log.Fatal(err)
			}
			log.Println(newWallet)
		}
	}
	// Тут прописаны возможные эндпоинты
	r := mux.NewRouter()
	r.HandleFunc("/api/send", handlers.Send).Methods("POST")
	r.HandleFunc("/api/transactions", handlers.GetLast).Methods("GET")
	r.HandleFunc("/api/wallet/{address}/balance", handlers.GetBalance).Methods("GET")
	// Запускаем http-сервер на порту 8080
	err = http.ListenAndServe("0.0.0.0:8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
