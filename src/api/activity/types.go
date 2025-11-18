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
	StartTime time.Time      `json:"start_time"`
	EndTime   time.Time      `json:"end_time"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ActivityParamsInput struct {
	ID uint `params:"id"`
}

type ActivityCreateInput struct {
	ProjectID uint      `validate:"required"`
	EditorID  uint      `validate:"required"`
	Language  string    `validate:"required"`
	FilePath  string    `validate:"required"`
	StartTime time.Time `validate:"required"`
	EndTime   time.Time `validate:"required,gtefield=StartTime"`
}

type ActivityUpdateInput struct {
	ProjectID uint      `validate:"required"`
	EditorID  uint      `validate:"required"`
	Language  string    `validate:"required"`
	FilePath  string    `validate:"required"`
	StartTime time.Time `validate:"required"`
	EndTime   time.Time `validate:"required,gtfield=StartTime"`
}
