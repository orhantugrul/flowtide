package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name       string     `gorm:"size:255;not null;uniqueIndex:idx_name_path"`
	Path       string     `gorm:"type:text;not null;uniqueIndex:idx_name_path"`
	Activities []Activity `gorm:"constraint:OnDelete:CASCADE;"`
}
