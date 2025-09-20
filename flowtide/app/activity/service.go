package activity

import (
	"context"

	"github.com/orhantugrul/flowtide/database"
	"gorm.io/gorm"
)

func GetActivities() ([]Activity, error) {
	return gorm.G[Activity](database.Database).
		Find(context.Background())
}

func GetActivity(id int) (Activity, error) {
	return gorm.G[Activity](database.Database).
		Where("id = ?", id).
		First(context.Background())
}

func CreateActivity(schema *ActivityCreateSchema) (*Activity, error) {
	activity := schema.ToModel()
	context := context.Background()

	err := gorm.G[Activity](database.Database).Create(context, &activity)
	if err != nil {
		return nil, err
	}

	return &activity, nil
}
