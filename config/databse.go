package config

import (
	"github.com/c0utin/historinha/helper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}

