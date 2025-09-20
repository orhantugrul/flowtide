package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

type Config struct {
	DatabasePath string
	LogLevel     logger.LogLevel
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
}

func Connect() error {
	config := Config{
		DatabasePath: func() string {
			homeDirectory, err := os.UserHomeDir()
			if err != nil {
				panic("user home directory not found")
			}

			flowtideDirectory := homeDirectory + "/.flowtide"
			if _, err := os.Stat(flowtideDirectory); os.IsNotExist(err) {
				if err := os.MkdirAll(flowtideDirectory, 0755); err != nil {
					return ".flowtide/flowtide.db"
				}
			}

			return flowtideDirectory + "/flowtide.db"
		}(),
		LogLevel:     logger.Info,
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		MaxLifetime:  time.Hour,
	}

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  config.LogLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	database, err := gorm.Open(sqlite.Open(config.DatabasePath), &gorm.Config{
		Logger: gormLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDatabase, err := database.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDatabase.SetMaxIdleConns(config.MaxIdleConns)
	sqlDatabase.SetMaxOpenConns(config.MaxOpenConns)
	sqlDatabase.SetConnMaxLifetime(config.MaxLifetime)

	if err := sqlDatabase.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	Database = database
	return nil
}

func Close() error {
	if Database == nil {
		return nil
	}

	sqlDB, err := Database.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return Database
}

func IsConnected() bool {
	if Database == nil {
		return false
	}

	sqlDB, err := Database.DB()
	if err != nil {
		return false
	}

	return sqlDB.Ping() == nil
}

func Transaction(fn func(*gorm.DB) error) error {
	if Database == nil {
		return fmt.Errorf("database connection not initialized")
	}

	return Database.Transaction(fn)
}
