package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Не вдалося підключитись до бази даних")
	}

	err = db.AutoMigrate(&Art{}, User{}, Admin{}, Basket{})
	if err != nil {
		panic("Migration error")
	}

	return db
}

func init() {
	DB = initDB()
}
