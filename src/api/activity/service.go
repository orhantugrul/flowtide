package activity

import (
	"context"

	"github.com/orhantugrul/flowtide/src/database"
	"github.com/orhantugrul/flowtide/src/database/model"
	"gorm.io/gorm"
)

func GetActivities(query *ActivityQueryInput) ([]model.Activity, error) {
	statement := gorm.G[model.Activity](database.Database)
	countStatement := gorm.G[model.Activity](database.Database)

	if !query.StartedAt.IsZero() {
		statement.Where("started_at >= ?", query.StartedAt)
		countStatement.Where("started_at >= ?", query.StartedAt)
	}

	if !query.EndedAt.IsZero() {
		statement.Where("ended_at <= ?", query.EndedAt)
		countStatement.Where("ended_at <= ?", query.EndedAt)
	}

	offset := int(query.Page * query.Size)
	limit := int(query.Size)
	return statement.Offset(offset).Limit(limit).Find(context.Background())
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
		StartedAt: body.StartedAt,
		EndedAt:   body.EndedAt,
	}

	err := gorm.G[model.Activity](database.Database).
		Create(context.Background(), &activity)
	return activity, err
}

func CreateActivities(body *[]ActivityCreateInput) ([]model.Activity, error) {
	activities := []model.Activity{}
	for _, item := range *body {
		activities = append(activities, model.Activity{
			ProjectID: item.ProjectID,
			EditorID:  item.EditorID,
			Language:  item.Language,
			FilePath:  item.FilePath,
			StartedAt: item.StartedAt,
			EndedAt:   item.EndedAt,
		})
	}

	err := gorm.G[[]model.Activity](database.Database).
		Create(context.Background(), &activities)
	return activities, err
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
		StartedAt: body.StartedAt,
		EndedAt:   body.EndedAt,
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
