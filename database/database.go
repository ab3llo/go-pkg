package database

import (
	"github.com/ab3llo/go-pkg/logger"
	"gofr.dev/pkg/gofr/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	log = logger.NewLogger()
)

func ConnectDB(cfg config.Config) {
	var err error

	DB, err = gorm.Open(postgres.Open(cfg.GetOrDefault("SUPABASE_DSN", "")))

	if err != nil {
		log.Error("Failed to connect to database")
	}

	log.Info("Database connection opened")
}
