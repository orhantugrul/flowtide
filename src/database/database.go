package database

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	config, err := NewConfig()
	if err != nil {
		return fmt.Errorf("failed to configurate the database: %w", err)
	}

	database, err := gorm.Open(sqlite.Open(config.Path), &gorm.Config{
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

	if err := migrate(sqlDatabase); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	Database = database
	return nil
}

func migrate(database *sql.DB) error {
	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	migrations := filepath.Join("src", "database", "migration")
	return goose.RunContext(context.Background(), "up", database, migrations)
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
