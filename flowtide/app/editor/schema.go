package editor

import (
	"time"

	"gorm.io/gorm"
)

type EditorSchema struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Version   string         `json:"version"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type EditorCreateSchema struct {
	Name    string `validate:"required"`
	Version string `validate:"required"`
}

func (editor *EditorCreateSchema) ToModel() Editor {
	return Editor{
		Name:    editor.Name,
		Version: editor.Version,
	}
}
