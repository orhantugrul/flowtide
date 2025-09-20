package activity

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	ProjectID uint
	EditorID  uint
	Language  string
	FilePath  string
	StartTime time.Time
	EndTime   time.Time
}

func (activity *Activity) toSchema() ActivitySchema {
	return ActivitySchema{
		ID:        activity.ID,
		ProjectID: activity.ProjectID,
		EditorID:  activity.EditorID,
		Language:  activity.Language,
		FilePath:  activity.FilePath,
		StartTime: activity.StartTime,
		EndTime:   activity.EndTime,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
		DeletedAt: activity.DeletedAt,
	}
}
