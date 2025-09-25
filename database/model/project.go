package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name       string     `gorm:"size:255;not null;index"`
	Path       string     `gorm:"type:text;not null"`
	Activities []Activity `gorm:"constraint:OnDelete:CASCADE;"`
}
