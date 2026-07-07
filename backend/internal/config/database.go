package config

import (
	"fmt"
	"log"
	"morng-dev/internal/adapters/persistence/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_NAME, config.DB_PORT, config.DB_SSLMODE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	runmigrate(db)
	log.Println("Database connected successfully")
	return db
}

func runmigrate(db *gorm.DB) {
	log.Println("Starting database migration...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Todo{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed successfully")
}
