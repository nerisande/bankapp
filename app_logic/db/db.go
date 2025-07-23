// Данный модуль содержит функцию для инициализации БД

package db

import (
	"bankapp/app_logic/models"
	"os"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Создаем объект БД, к которому можно будет в дальнейшем обращаться как к db.DB
var DB *gorm.DB

func Init_db() (*gorm.DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "database.db"
	}
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&models.Wallet{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&models.Transaction{}); err != nil {
		return DB, err
	}
	return DB, nil
}
