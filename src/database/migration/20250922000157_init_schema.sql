-- +goose Up
CREATE TABLE IF NOT EXISTS projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    path TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    UNIQUE(name, path)
);

CREATE TABLE IF NOT EXISTS editors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    version VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    UNIQUE(name, version)
);

CREATE TABLE IF NOT EXISTS activities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    project_id INTEGER NOT NULL,
    editor_id INTEGER NOT NULL,
    language VARCHAR(100) NOT NULL,
    file_path TEXT NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (editor_id) REFERENCES editors(id) ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_projects_name_path ON projects(name, path);
CREATE INDEX IF NOT EXISTS idx_projects_deleted_at ON projects(deleted_at);

CREATE INDEX IF NOT EXISTS idx_editors_name_version ON editors(name, version);
CREATE INDEX IF NOT EXISTS idx_editors_deleted_at ON editors(deleted_at);

CREATE INDEX IF NOT EXISTS idx_activities_project_id ON activities(project_id);
CREATE INDEX IF NOT EXISTS idx_activities_editor_id ON activities(editor_id);
CREATE INDEX IF NOT EXISTS idx_activities_language ON activities(language);
CREATE INDEX IF NOT EXISTS idx_activities_start_time ON activities(start_time);
CREATE INDEX IF NOT EXISTS idx_activities_end_time ON activities(end_time);
CREATE INDEX IF NOT EXISTS idx_activities_deleted_at ON activities(deleted_at);

CREATE INDEX IF NOT EXISTS idx_activities_project_time ON activities(project_id, start_time, end_time);
CREATE INDEX IF NOT EXISTS idx_activities_editor_time ON activities(editor_id, start_time, end_time);


-- +goose Down
DROP INDEX IF EXISTS idx_activities_editor_time;
DROP INDEX IF EXISTS idx_activities_project_time;

DROP INDEX IF EXISTS idx_activities_deleted_at;
DROP INDEX IF EXISTS idx_activities_end_time;
DROP INDEX IF EXISTS idx_activities_start_time;
DROP INDEX IF EXISTS idx_activities_language;
DROP INDEX IF EXISTS idx_activities_editor_id;
DROP INDEX IF EXISTS idx_activities_project_id;

DROP INDEX IF EXISTS idx_editors_deleted_at;
DROP INDEX IF EXISTS idx_editors_name_version;

DROP INDEX IF EXISTS idx_projects_deleted_at;
DROP INDEX IF EXISTS idx_projects_name_path;

DROP TABLE IF EXISTS activities;
DROP TABLE IF EXISTS editors;
DROP TABLE IF EXISTS projects;

