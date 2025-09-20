-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL
);

-- Create projects table
CREATE TABLE IF NOT EXISTS projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    path TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create editors table
CREATE TABLE IF NOT EXISTS editors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE,
    version VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL
);

-- Create activities table
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

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);

CREATE INDEX IF NOT EXISTS idx_projects_user_id ON projects(user_id);
CREATE INDEX IF NOT EXISTS idx_projects_name ON projects(name);
CREATE INDEX IF NOT EXISTS idx_projects_deleted_at ON projects(deleted_at);

CREATE INDEX IF NOT EXISTS idx_editors_name ON editors(name);
CREATE INDEX IF NOT EXISTS idx_editors_deleted_at ON editors(deleted_at);

CREATE INDEX IF NOT EXISTS idx_activities_project_id ON activities(project_id);
CREATE INDEX IF NOT EXISTS idx_activities_editor_id ON activities(editor_id);
CREATE INDEX IF NOT EXISTS idx_activities_language ON activities(language);
CREATE INDEX IF NOT EXISTS idx_activities_start_time ON activities(start_time);
CREATE INDEX IF NOT EXISTS idx_activities_end_time ON activities(end_time);
CREATE INDEX IF NOT EXISTS idx_activities_deleted_at ON activities(deleted_at);

-- Create composite indexes for common queries
CREATE INDEX IF NOT EXISTS idx_activities_project_time ON activities(project_id, start_time, end_time);
CREATE INDEX IF NOT EXISTS idx_activities_editor_time ON activities(editor_id, start_time, end_time);
