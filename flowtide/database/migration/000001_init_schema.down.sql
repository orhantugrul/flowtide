-- Drop indexes first
DROP INDEX IF EXISTS idx_activities_editor_time;
DROP INDEX IF EXISTS idx_activities_project_time;

DROP INDEX IF EXISTS idx_activities_deleted_at;
DROP INDEX IF EXISTS idx_activities_end_time;
DROP INDEX IF EXISTS idx_activities_start_time;
DROP INDEX IF EXISTS idx_activities_language;
DROP INDEX IF EXISTS idx_activities_editor_id;
DROP INDEX IF EXISTS idx_activities_project_id;

DROP INDEX IF EXISTS idx_editors_deleted_at;
DROP INDEX IF EXISTS idx_editors_name;

DROP INDEX IF EXISTS idx_projects_deleted_at;
DROP INDEX IF EXISTS idx_projects_name;
DROP INDEX IF EXISTS idx_projects_user_id;

DROP INDEX IF EXISTS idx_users_deleted_at;
DROP INDEX IF EXISTS idx_users_username;

-- Drop tables in reverse order (respecting foreign key dependencies)
DROP TABLE IF EXISTS activities;
DROP TABLE IF EXISTS editors;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS users;
