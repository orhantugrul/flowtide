import { Editor, EditorCreateInput, EditorQueryInput } from "src/types/editor";
import { fetcher, FetchResult } from "src/utils/fetcher";

export async function getEditors(
  query: EditorQueryInput
): Promise<FetchResult<Editor[]>> {
  const params = new URLSearchParams(Object.entries(query)).toString();
  return await fetcher<Editor[]>(`/api/editors?${params}`);
}

export async function createEditor(
  body: EditorCreateInput
): Promise<FetchResult<Editor>> {
  return await fetcher<Editor>(`/api/editors`, {
    method: "POST",
    body: JSON.stringify(body),
  });
}
