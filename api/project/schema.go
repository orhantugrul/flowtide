package project

import (
	"time"

	"gorm.io/gorm"
)

type ProjectParamsSchema struct {
	ID uint `params:"id"`
}

type ProjectSchema struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ProjectCreateSchema struct {
	UserID uint   `validate:"required"`
	Name   string `validate:"required"`
	Path   string `validate:"required"`
}

type ProjectUpdateSchema struct {
	UserID uint   `validate:"required"`
	Name   string `validate:"required"`
	Path   string `validate:"required"`
}
