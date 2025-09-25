package project

import (
	"context"

	"github.com/orhantugrul/flowtide/database"
	"github.com/orhantugrul/flowtide/database/model"
	"gorm.io/gorm"
)

func GetProjects() ([]model.Project, error) {
	return gorm.G[model.Project](database.Database).Find(context.Background())
}

func GetProject(id uint) (model.Project, error) {
	return gorm.G[model.Project](database.Database).
		Where("id = ?", id).
		First(context.Background())
}

func GetProjectActivities(id uint) ([]model.Activity, error) {
	return gorm.G[model.Activity](database.Database).
		Where("project_id = ?", id).
		Find(context.Background())
}

func CreateProject(body *ProjectCreateSchema) (model.Project, error) {
	project := model.Project{
		Name: body.Name,
		Path: body.Path,
	}

	err := gorm.G[model.Project](database.Database).
		Create(context.Background(), &project)

	if err != nil {
		return model.Project{}, err
	}

	return GetProject(project.ID)
}

func UpdateProject(id uint, body *ProjectUpdateSchema) (model.Project, error) {
	project := model.Project{
		Name: body.Name,
		Path: body.Path,
	}

	_, err := gorm.G[model.Project](database.Database).
		Where("id = ?", id).
		Updates(context.Background(), project)

	if err != nil {
		return model.Project{}, err
	}

	return GetProject(id)
}

func DeleteProject(id uint) error {
	_, err := gorm.G[model.Project](database.Database).
		Where("id = ?", id).
		Delete(context.Background())
	return err
}
