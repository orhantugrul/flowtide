import {
  Project,
  ProjectCreateInput,
  ProjectQueryInput,
} from "src/types/project";
import { fetcher, FetchResult } from "src/utils/fetcher";

export async function getProjects(
  query: ProjectQueryInput
): Promise<FetchResult<Project[]>> {
  const params = new URLSearchParams(Object.entries(query)).toString();
  return await fetcher<Project[]>(`/api/projects?${params}`);
}

export async function createProject(
  body: ProjectCreateInput
): Promise<FetchResult<Project>> {
  return await fetcher<Project>(`/api/projects`, {
    method: "POST",
    body: JSON.stringify(body),
  });
}
