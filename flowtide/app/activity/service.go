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

	err := gorm.G[Activity](database.Database).
		Create(context.Background(), &activity)

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func UpdateActivity(id int, schema *ActivityUpdateSchema) (*Activity, error) {
	activity := schema.ToModel()

	_, err := gorm.G[Activity](database.Database).
		Where("id = ?", id).
		Updates(context.Background(), activity)

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func DeleteActivity(id int) error {
	_, err := gorm.G[Activity](database.Database).
		Where("id = ?", id).
		Delete(context.Background())

	if err != nil {
		return err
	}

	return nil
}
