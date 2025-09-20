package editor

import (
	"context"

	"github.com/orhantugrul/flowtide/database"
	"gorm.io/gorm"
)

func GetEditors() ([]Editor, error) {
	return gorm.G[Editor](database.Database).
		Find(context.Background())
}

func GetEditor(id int) (Editor, error) {
	return gorm.G[Editor](database.Database).
		Where("id = ?", id).
		First(context.Background())
}

func CreateEditor(schema *EditorCreateSchema) (*Editor, error) {
	editor := schema.ToModel()

	err := gorm.G[Editor](database.Database).
		Create(context.Background(), &editor)

	if err != nil {
		return nil, err
	}

	return &editor, nil
}

func UpdateEditor(id int, schema *EditorUpdateSchema) (*Editor, error) {
	editor := schema.ToModel()

	_, err := gorm.G[Editor](database.Database).
		Where("id = ?", id).
		Updates(context.Background(), editor)

	if err != nil {
		return nil, err
	}

	return &editor, nil
}

func DeleteEditor(id int) error {
	_, err := gorm.G[Editor](database.Database).
		Where("id = ?", id).
		Delete(context.Background())

	if err != nil {
		return err
	}

	return nil
}
