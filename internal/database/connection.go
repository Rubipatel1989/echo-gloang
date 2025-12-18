package database

import (
	"fmt"
	"log"

	"echo-golang/internal/config"
	"echo-golang/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	var err error

	dsn := config.AppConfig.GetDSN()

	// Configure GORM logger based on environment
	var gormLogger logger.Interface
	if config.AppConfig.Env == "production" {
		gormLogger = logger.Default.LogMode(logger.Silent)
	} else {
		gormLogger = logger.Default.LogMode(logger.Info)
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")

	// Auto migrate models
	if err := AutoMigrate(); err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}

	return nil
}

func AutoMigrate() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Organization{},
	)
}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

