package project

import (
	"time"

	"gorm.io/gorm"
)

type ProjectSchema struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"user_id"`
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

func (project *ProjectCreateSchema) ToModel() Project {
	return Project{
		UserID: project.UserID,
		Name:   project.Name,
		Path:   project.Path,
	}
}

type ProjectUpdateSchema struct {
	UserID uint   `validate:"required"`
	Name   string `validate:"required"`
	Path   string `validate:"required"`
}

func (project *ProjectUpdateSchema) ToModel() Project {
	return Project{
		UserID: project.UserID,
		Name:   project.Name,
		Path:   project.Path,
	}
}
