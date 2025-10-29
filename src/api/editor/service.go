package editor

import (
	"context"

	"github.com/orhantugrul/flowtide/src/database"
	"github.com/orhantugrul/flowtide/src/database/model"
	"gorm.io/gorm"
)

func GetEditors(query *EditorQueryInput) ([]model.Editor, error) {
	return gorm.G[model.Editor](database.Database).
		Where(&model.Editor{Name: query.Name, Version: query.Version}).
		Find(context.Background())

}

func GetEditor(id uint) (model.Editor, error) {
	return gorm.G[model.Editor](database.Database).
		Where("id = ?", id).
		First(context.Background())
}

func CreateEditor(body *EditorCreateInput) (model.Editor, error) {
	editor := model.Editor{
		Name:    body.Name,
		Version: body.Version,
	}

	err := gorm.G[model.Editor](database.Database).
		Create(context.Background(), &editor)
	return editor, err
}

func UpdateEditor(id uint, body *EditorUpdateInput) (model.Editor, error) {
	editor := model.Editor{
		Name:    body.Name,
		Version: body.Version,
	}

	_, err := gorm.G[model.Editor](database.Database).
		Where("id = ?", id).
		Updates(context.Background(), editor)

	if err != nil {
		return model.Editor{}, err
	}

	return GetEditor(id)
}

func DeleteEditor(id uint) error {
	_, err := gorm.G[model.Editor](database.Database).
		Where("id = ?", id).
		Delete(context.Background())
	return err
}
