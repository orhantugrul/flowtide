package project

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID uint
	Name   string
	Path   string
}

func (project *Project) ToSchema() ProjectSchema {
	return ProjectSchema{
		ID:        project.ID,
		UserID:    project.UserID,
		Name:      project.Name,
		Path:      project.Path,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
		DeletedAt: project.DeletedAt,
	}
}
