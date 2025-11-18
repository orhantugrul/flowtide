import type { Activity, ActivityCreateInput } from "../types/activity";
import { type FetchResult, fetcher } from "../utils/fetcher";

export async function getActivities(): Promise<FetchResult<Activity[]>> {
  return await fetcher<Activity[]>("/api/activities");
}

export async function createActivity(
  body: ActivityCreateInput,
): Promise<FetchResult<Activity>> {
  return await fetcher<Activity>("/api/activities", {
    method: "POST",
    body: JSON.stringify(body),
  });
}

export async function createActivities(
  body: ActivityCreateInput[],
): Promise<FetchResult<Activity[]>> {
  return await fetcher<Activity[]>("/api/activities/batch", {
    method: "POST",
    body: JSON.stringify(body),
  });
}
