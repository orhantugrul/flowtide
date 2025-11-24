package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

type Config struct {
	Path         string
	Logger       logger.Config
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
}

func NewConfig() (*Config, error) {
	path, err := getDatabasePath()
	if err != nil {
		return nil, fmt.Errorf("failed to get database path: %w", err)
	}

	return &Config{
		Path:         path,
		MaxIdleConns: 4,
		MaxOpenConns: 8,
		MaxLifetime:  time.Hour,
	}, nil
}

func getDatabasePath() (string, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home directory not found: %w", err)
	}

	flowtideDirectory := homeDirectory + "/.flowtide"
	if _, err := os.Stat(flowtideDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(flowtideDirectory, 0755); err != nil {
			return "", fmt.Errorf("failed to create flowtide directory: %w", err)
		}
	}

	return flowtideDirectory + "/flowtide.db", nil
}
