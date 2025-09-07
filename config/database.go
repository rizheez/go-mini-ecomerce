package config

import (
	"fmt"

	"mini-ecommerce/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func Connect(cfg *Config) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DB.Host, cfg.DB.User, cfg.DB.Pass, cfg.DB.DBName, cfg.DB.Port, cfg.DB.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal(err, "[ErrDatabase-1]Failed to connect to database")
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error(err, "[ErrDatabase-2]Failed to get database instance")
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		logger.Error(err, "[ErrDatabase-3]Failed to ping database")
		return nil, err
	}

	logger.Info("Successfully connected to database")
	return &Postgres{DB: db}, nil
}

func (p *Postgres) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		logger.Error(err, "[ErrDatabase-4]Failed to get database instance for closing")
		return err
	}

	if err := sqlDB.Close(); err != nil {
		logger.Error(err, "[ErrDatabase-5]Failed to close database connection")
		return err
	}

	logger.Info("Database connection closed")
	return nil
}
