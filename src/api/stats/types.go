package stats

import "time"

type StatsRangeQueryInput struct {
	start string `query:"start"`
	end   string `query:"end"`
}

type StatsDateRange struct {
	Start time.Time
	End   time.Time
}

type StatsHoursOutput struct {
	TotalHours float64   `json:"total_hours"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

type StatsDailyActivityOutput struct {
	Date     time.Time `json:"date"`
	Hours    float64   `json:"hours"`
	Sessions int       `json:"sessions"`
}

type StatsLanguageUsageOutput struct {
	Language   string  `json:"language"`
	Hours      float64 `json:"hours"`
	Percentage float64 `json:"percentage"`
	Sessions   int     `json:"sessions"`
}

type StatsEditorUsageOutput struct {
	EditorID   uint    `json:"editor_id"`
	Name       string  `json:"name"`
	Version    string  `json:"version"`
	Hours      float64 `json:"hours"`
	Percentage float64 `json:"percentage"`
	Sessions   int     `json:"sessions"`
}

type StatsProjectsOutput struct {
	Total int                       `json:"total"`
	Items []StatsProjectUsageOutput `json:"items"`
}

type StatsProjectUsageOutput struct {
	ProjectID uint      `json:"project_id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Hours     float64   `json:"hours"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
