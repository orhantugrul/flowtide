package activity

import (
	"context"

	"github.com/orhantugrul/flowtide/database"
	"github.com/orhantugrul/flowtide/database/model"
	"gorm.io/gorm"
)

func GetActivities() ([]model.Activity, error) {
	return gorm.G[model.Activity](database.Database).Find(context.Background())
}

func GetActivity(id uint) (model.Activity, error) {
	return gorm.G[model.Activity](database.Database).
		Where("id = ?", id).
		First(context.Background())
}

func CreateActivity(body *ActivityCreateInput) (model.Activity, error) {
	activity := model.Activity{
		ProjectID: body.ProjectID,
		EditorID:  body.EditorID,
		Language:  body.Language,
		FilePath:  body.FilePath,
		StartTime: body.StartTime,
		EndTime:   body.EndTime,
	}

	err := gorm.G[model.Activity](database.Database).
		Create(context.Background(), &activity)
	return activity, err
}

func UpdateActivity(
	id uint,
	body *ActivityUpdateInput,
) (model.Activity, error) {
	activity := model.Activity{
		ProjectID: body.ProjectID,
		EditorID:  body.EditorID,
		Language:  body.Language,
		FilePath:  body.FilePath,
		StartTime: body.StartTime,
		EndTime:   body.EndTime,
	}

	_, err := gorm.G[model.Activity](database.Database).
		Where("id = ?", id).
		Updates(context.Background(), activity)

	if err != nil {
		return model.Activity{}, err
	}

	return GetActivity(id)
}

func DeleteActivity(id uint) error {
	_, err := gorm.G[model.Activity](database.Database).
		Where("id = ?", id).
		Delete(context.Background())
	return err
}
