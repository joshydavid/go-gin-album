package db

import (
	"fmt"
	"go-gin-album/internal/config"
	"go-gin-album/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dbConfig *config.DBConfig) (*gorm.DB, error) {
	log.Println("🚀 Attempting to connect to PostgreSQL")
	db, err := gorm.Open(postgres.Open(dbConfig.DSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	log.Println("✅ Successfully connected to PostgreSQL")

	if err := db.AutoMigrate(&model.Album{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate PostgreSQL schema: %w", err)
	}

	return db, nil
}
