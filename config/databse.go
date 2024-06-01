package config

import (
	"github.com/c0utin/historinha/helper"
	"github.com/c0utin/historinha/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		helper.ErrorPanic(err)
	}
}

func DatabaseConnection() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	helper.ErrorPanic(err)

	// Auto migrate
	db.AutoMigrate(&model.Users{}, &model.Historys{})

	return db
}

