package project

import (
	"time"

	"gorm.io/gorm"
)

type ProjectOutput struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ProjectParamsInput struct {
	ID uint `params:"id"`
}

type ProjectQueryInput struct {
	Name string `query:"name"`
	Path string `query:"path"`
}

type ProjectCreateInput struct {
	Name string `validate:"required"`
	Path string `validate:"required"`
}

type ProjectUpdateInput struct {
	Name string `validate:"required"`
	Path string `validate:"required"`
}
