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
	context := context.Background()

	err := gorm.G[Editor](database.Database).Create(context, &editor)
	if err != nil {
		return nil, err
	}

	return &editor, nil
}
