import * as vscode from "vscode";
import { createEditor, getEditors } from "../api/editors";
import { createProject, getProjects } from "../api/projects";
import { Editor } from "../types/editor";
import { Project } from "../types/project";

/**
 * Gets the current editor instance for the running VS Code environment.
 * If the editor is not found in the backend, it will be created.
 *
 * @returns {Promise<Editor>} A promise that resolves to the current Editor object.
 * @throws {Error} If the editor cannot be created or retrieved.
 */
export async function getCurrentEditor(): Promise<Editor> {
  const name = vscode.env.appName.toLowerCase();
  const version = vscode.version;

  const { data: editors } = await getEditors({ name, version });
  if (editors?.[0]) {
    return editors[0];
  }

  const { data: editor, error } = await createEditor({ name, version });
  if (error) {
    throw new Error(error.message);
  }

  return editor!;
}

/**
 * Gets the current project associated with the open workspace.
 * If the project is not found in the backend, it will be created.
 *
 * @returns {Promise<Project>} A promise that resolves to the current Project object.
 * @throws {Error} If the project cannot be created or retrieved.
 */
export async function getCurrentProject(): Promise<Project> {
  const folder = vscode.workspace.workspaceFolders?.[0];
  if (!folder) {
    throw new Error("No workspace folder found");
  }

  const { name, uri } = folder;
  const path = uri.fsPath;

  const { data: projects } = await getProjects({ name, path });
  if (projects?.[0]) {
    return projects[0];
  }

  const { data: project, error } = await createProject({ name, path });
  if (error) {
    throw new Error(error.message);
  }

  return project!;
}
