import * as vscode from "vscode";

interface FetchError {
  code: string;
  message: string;
}

export async function fetcher<T>(
  path: string,
  init?: RequestInit
): Promise<[T, FetchError]> {
  const url = vscode.workspace.getConfiguration("flowtide").get<string>("url");

  const response = await fetch(`${url}${path}`, {
    headers: {
      "Content-Type": "application/json",
      ...(init?.headers || {}),
    },
    ...init,
  });

  if (!response.ok) {
    const body = (await response.json()) as { message?: string; code?: string };
    return [
      {} as T,
      {
        code: body?.code || response.status.toString(),
        message: body?.message || "Unknown error occurred",
      },
    ];
  }

  if (response.status === 204) {
    return [{} as T, {} as FetchError];
  }

  const body = await response.json();
  return [body as T, {} as FetchError];
}
