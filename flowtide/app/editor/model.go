package editor

import (
	"github.com/orhantugrul/flowtide/app/activity"
	"gorm.io/gorm"
)

type Editor struct {
	gorm.Model
	Name       string
	Version    string
	Activities []activity.Activity
}

func (editor *Editor) ToSchema() EditorSchema {
	return EditorSchema{
		ID:        editor.ID,
		Name:      editor.Name,
		Version:   editor.Version,
		CreatedAt: editor.CreatedAt,
		UpdatedAt: editor.UpdatedAt,
		DeletedAt: editor.DeletedAt,
	}
}
