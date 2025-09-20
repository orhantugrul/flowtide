package project

import (
	"context"

	"github.com/orhantugrul/flowtide/app/activity"
	"github.com/orhantugrul/flowtide/database"
	"gorm.io/gorm"
)

func GetProjects() ([]Project, error) {
	return gorm.G[Project](database.Database).Find(context.Background())
}

func GetProject(id int) (Project, error) {
	return gorm.G[Project](database.Database).
		Where("id = ?", id).
		First(context.Background())
}

func GetProjectActivities(id int) ([]activity.Activity, error) {
	return gorm.G[activity.Activity](database.Database).
		Where("project_id = ?", id).
		Find(context.Background())
}

func CreateProject(schema *ProjectCreateSchema) (*Project, error) {
	project := schema.ToModel()

	err := gorm.G[Project](database.Database).
		Create(context.Background(), &project)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func UpdateProject(id int, schema *ProjectUpdateSchema) (*Project, error) {
	project := schema.ToModel()

	_, err := gorm.G[Project](database.Database).
		Where("id = ?", id).
		Updates(context.Background(), project)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func DeleteProject(id int) error {
	_, err := gorm.G[Project](database.Database).
		Where("id = ?", id).
		Delete(context.Background())

	if err != nil {
		return err
	}

	return nil
}
