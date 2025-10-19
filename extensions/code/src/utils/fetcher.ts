import * as vscode from "vscode";

export interface FetchResult<T> {
  data?: T;
  error?: FetchError;
}

export interface FetchError {
  code: string;
  message: string;
}

export async function fetcher<T>(
  path: string,
  init?: RequestInit
): Promise<FetchResult<T>> {
  const url = vscode.workspace.getConfiguration("flowtide").get<string>("url");

  const response = await fetch(`${url}${path}`, {
    headers: {
      "Content-Type": "application/json",
      ...(init?.headers || {}),
    },
    ...init,
  });

  if (!response.ok) {
    const body = await response.json();
    const error = body as FetchError;

    return {
      data: undefined,
      error: {
        code: error?.code ?? response.status.toString(),
        message: error?.message ?? "Unknown error occurred",
      },
    };
  }

  const body = (await response.json().catch(() => undefined)) as T;
  return { data: body, error: undefined };
}
