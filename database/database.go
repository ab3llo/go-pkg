package database

import (
	"github.com/ab3llo/go-pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) *gorm.DB {
	logger := logger.NewLogger()

	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		logger.Error("Failed to connect to database")
	}

	logger.Info("Database connection opened")
	return DB
}
