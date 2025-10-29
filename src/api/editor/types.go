package editor

import (
	"time"

	"gorm.io/gorm"
)

type EditorOutput struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Version   string         `json:"version"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type EditorParamsInput struct {
	ID uint `params:"id"`
}

type EditorQueryInput struct {
	Name    string `query:"name"`
	Version string `query:"version"`
}

type EditorCreateInput struct {
	Name    string `validate:"required"`
	Version string `validate:"required"`
}

type EditorUpdateInput struct {
	Name    string `validate:"required"`
	Version string `validate:"required"`
}
