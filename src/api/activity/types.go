package activity

import (
	"time"

	"gorm.io/gorm"
)

type ActivityOutput struct {
	ID        uint           `json:"id"`
	ProjectID uint           `json:"project_id"`
	EditorID  uint           `json:"editor_id"`
	Language  string         `json:"language"`
	FilePath  string         `json:"file_path"`
	StartedAt time.Time      `json:"started_at"`
	EndedAt   time.Time      `json:"ended_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ActivityParamsInput struct {
	ID uint `params:"id"`
}

type ActivityQueryInput struct {
	Page      uint      `query:"page"`
	Size      uint      `query:"size"`
	StartedAt time.Time `query:"started_at"`
	EndedAt   time.Time `query:"ended_at"`
}

type ActivityCreateInput struct {
	ProjectID uint      `validate:"required"`
	EditorID  uint      `validate:"required"`
	Language  string    `validate:"required"`
	FilePath  string    `validate:"required"`
	StartedAt time.Time `validate:"required"`
	EndedAt   time.Time `validate:"required,gtefield=StartedAt"`
}

type ActivityUpdateInput struct {
	ProjectID uint      `validate:"required"`
	EditorID  uint      `validate:"required"`
	Language  string    `validate:"required"`
	FilePath  string    `validate:"required"`
	StartedAt time.Time `validate:"required"`
	EndedAt   time.Time `validate:"required,gtfield=StartedAt"`
}
