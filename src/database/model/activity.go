package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	ProjectID uint      `gorm:"not null;index"`
	EditorID  uint      `gorm:"not null;index"`
	Language  string    `gorm:"size:100;not null;index"`
	FilePath  string    `gorm:"type:text;not null"`
	StartTime time.Time `gorm:"not null;index"`
	EndTime   time.Time `gorm:"not null;index"`
	Project   Project   `gorm:"constraint:OnDelete:CASCADE;"`
	Editor    Editor    `gorm:"constraint:OnDelete:RESTRICT;"`
}
