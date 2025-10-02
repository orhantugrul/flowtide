import { Activity, ActivityCreateInput } from "../../types/activity";
import {
  Editor,
  EditorCreateInput,
  EditorQueryInput,
} from "../../types/editor";
import { Project, ProjectCreateInput } from "../../types/project";
import { fetcher } from "../../utils/fetcher";

export async function getEditors(query: EditorQueryInput): Promise<Editor[]> {
  const [editors, error] = await fetcher<Editor[]>(
    `/api/editors?name=${query.name}&version=${query.version}`
  );

  if (error.code) {
    throw new Error(error.message);
  }

  return editors;
}

export async function createEditor(
  payload: EditorCreateInput
): Promise<Editor> {
  const [created, error] = await fetcher<Editor>("/api/editors", {
    method: "POST",
    body: JSON.stringify(payload),
  });

  if (error.code) {
    throw new Error(error.message);
  }

  return created;
}

export async function createProject(
  payload: ProjectCreateInput
): Promise<Project> {
  const [project, error] = await fetcher<Project>("/api/projects", {
    method: "POST",
    body: JSON.stringify(payload),
  });

  if (error.code) {
    throw new Error(error.message);
  }

  return project;
}

export async function createActivity(
  payload: ActivityCreateInput
): Promise<Activity> {
  const [activity, error] = await fetcher<Activity>("/api/activities", {
    method: "POST",
    body: JSON.stringify(payload),
  });

  if (error.code) {
    throw new Error(error.message);
  }

  return activity;
}

export async function createActivities(
  payload: ActivityCreateInput[]
): Promise<Activity[]> {
  const [activities, error] = await fetcher<Activity[]>(
    "/api/activities/batch",
    {
      method: "POST",
      body: JSON.stringify(payload),
    }
  );

  if (error.code) {
    throw new Error(error.message);
  }

  return activities;
}
