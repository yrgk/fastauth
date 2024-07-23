package storage

import (
	"fastauth/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(config config.Config) *gorm.DB {
	// Getting database
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrating models
	db.AutoMigrate(
		User{},
	)

	return db
}