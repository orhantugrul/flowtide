package project

import (
	"os/user"

	"github.com/orhantugrul/flowtide/app/activity"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID     uint
	Name       string
	Path       string
	User       user.User
	Activities []activity.Activity
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
