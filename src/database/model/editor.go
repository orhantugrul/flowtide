package model

import (
	"gorm.io/gorm"
)

type Editor struct {
	gorm.Model
	Name       string     `gorm:"size:255;not null;uniqueIndex:idx_name_version"`
	Version    string     `gorm:"size:100;not null;uniqueIndex:idx_name_version"`
	Activities []Activity `gorm:"constraint:OnDelete:RESTRICT;"`
}
