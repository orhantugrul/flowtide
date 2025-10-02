import * as vscode from "vscode";
import { ActivityCreateInput } from "../../types/activity";
import { Editor } from "../../types/editor";
import { Project } from "../../types/project";
import { EXTENSIONS } from "../../utils/file";
import {
  createActivities,
  createEditor,
  createProject,
  getEditors,
} from "./api";

export interface TrackingSession {
  project: Project;
  editor: Editor;
  startTime: Date;
  file?: string;
  language?: string;
}

export class Tracker {
  private static instance: Tracker;
  private tracking = false;
  private session: TrackingSession | null = null;
  private statusBarItem: vscode.StatusBarItem;
  private flushTimer: NodeJS.Timeout | null = null;
  private activityBuffer: ActivityCreateInput[] = [];
  private disposables: vscode.Disposable[] = [];

  private constructor() {
    this.statusBarItem = vscode.window.createStatusBarItem(
      vscode.StatusBarAlignment.Right,
      100
    );
    this.statusBarItem.command = "flowtide.toggle";
    this.updateStatusBar();
  }

  public static getInstance(): Tracker {
    if (!Tracker.instance) {
      Tracker.instance = new Tracker();
    }
    return Tracker.instance;
  }

  public async startTracking(): Promise<void> {
    if (this.tracking) {
      return;
    }

    try {
      const editor = await this.getOrCreateEditor();
      const project = await this.getOrCreateProject();

      this.session = {
        project,
        editor,
        startTime: new Date(),
      };

      this.tracking = true;
      this.setupEventListeners();
      this.startFlushTimer();
      this.updateStatusBar();

      vscode.window.showInformationMessage("🎯 Flowtide tracking started!");

      // Set context for command palette
      vscode.commands.executeCommand("setContext", "flowtide.tracking", true);
    } catch (error) {
      vscode.window.showErrorMessage(
        `Failed to start tracking: ${
          error instanceof Error ? error.message : "Unknown error"
        }`
      );
    }
  }

  public async stopTracking(): Promise<void> {
    if (!this.tracking) {
      return;
    }

    try {
      // Flush any remaining activities
      await this.flushActivities();

      this.tracking = false;
      this.session = null;
      this.cleanup();
      this.updateStatusBar();

      vscode.window.showInformationMessage("⏹️ Flowtide tracking stopped!");

      // Set context for command palette
      vscode.commands.executeCommand("setContext", "flowtide.tracking", false);
    } catch (error) {
      vscode.window.showErrorMessage(
        `Failed to stop tracking: ${
          error instanceof Error ? error.message : "Unknown error"
        }`
      );
    }
  }

  public async toggleTracking(): Promise<void> {
    if (this.tracking) {
      await this.stopTracking();
    } else {
      await this.startTracking();
    }
  }

  private async getOrCreateEditor(): Promise<Editor> {
    const name = vscode.env.appName.toLowerCase();
    const version = vscode.version;

    const editors = await getEditors({ name, version });
    if (editors.length > 0) {
      return editors[0];
    }

    const editor = await createEditor({ name, version });
    return editor;
  }

  private async getOrCreateProject(): Promise<Project> {
    const workspaceFolder = vscode.workspace.workspaceFolders?.[0];
    if (!workspaceFolder) {
      throw new Error("No workspace folder found");
    }

    return await createProject({
      name: workspaceFolder.name,
      path: workspaceFolder.uri.fsPath,
    });
  }

  private setupEventListeners(): void {
    const onDidChangeActiveTextEditor =
      vscode.window.onDidChangeActiveTextEditor((editor) => {
        if (editor && this.tracking) {
          this.trackFileActivity(editor);
        }
      });

    const onDidChangeTextDocument = vscode.workspace.onDidChangeTextDocument(
      (event) => {
        if (
          this.tracking &&
          event.document === vscode.window.activeTextEditor?.document
        ) {
          this.trackFileActivity(vscode.window.activeTextEditor!);
        }
      }
    );

    // Track when user becomes idle/active
    const onDidChangeWindowState = vscode.window.onDidChangeWindowState(
      (state) => {
        if (this.tracking && state.focused) {
          // User became active, track current file
          const editor = vscode.window.activeTextEditor;
          if (editor) {
            this.trackFileActivity(editor);
          }
        }
      }
    );

    this.disposables.push(
      onDidChangeActiveTextEditor,
      onDidChangeTextDocument,
      onDidChangeWindowState
    );
  }

  private trackFileActivity(editor: vscode.TextEditor): void {
    if (!this.session) {
      return;
    }

    const filePath = editor.document.uri.fsPath;
    const language = this.detectLanguage(filePath);
    const now = new Date();

    // If we're tracking the same file, update the end time
    if (this.session.file === filePath) {
      // Update the last activity in buffer
      const lastActivity = this.activityBuffer[this.activityBuffer.length - 1];
      if (lastActivity && lastActivity.filePath === filePath) {
        lastActivity.endTime = now;
      }
    } else {
      // New file, create new activity
      const activity: ActivityCreateInput = {
        projectId: this.session.project.id!,
        editorId: this.session.editor.id!,
        language,
        filePath,
        startTime: now,
        endTime: now,
      };

      this.activityBuffer.push(activity);
      this.session.file = filePath;
      this.session.language = language;
    }
  }

  private detectLanguage(filePath: string): string {
    const lastIndex = filePath.lastIndexOf(".");
    const extension = filePath.substring(lastIndex);
    return EXTENSIONS[extension as keyof typeof EXTENSIONS] ?? "Unknown";
  }

  private startFlushTimer(): void {
    this.flushTimer = setInterval(this.flushActivities, 30000); // 30 seconds
  }

  private async flushActivities(): Promise<void> {
    if (this.activityBuffer.length === 0) {
      return;
    }

    try {
      await createActivities([...this.activityBuffer]);
      this.activityBuffer = [];
    } catch (error) {
      console.error("Failed to flush activities:", error);
      // Keep activities in buffer for retry
    }
  }

  private updateStatusBar(): void {
    if (this.tracking) {
      this.statusBarItem.text = "$(pulse) Flowtide";
      this.statusBarItem.tooltip = "Click to stop activity tracking";
      this.statusBarItem.backgroundColor = new vscode.ThemeColor(
        "statusBarItem.prominentBackground"
      );
    } else {
      this.statusBarItem.text = "$(circle-outline) Flowtide";
      this.statusBarItem.tooltip = "Click to start activity tracking";
      this.statusBarItem.backgroundColor = undefined;
    }
    this.statusBarItem.show();
  }

  private cleanup(): void {
    this.disposables.forEach((disposable) => disposable.dispose());
    this.disposables = [];

    if (this.flushTimer) {
      clearInterval(this.flushTimer);
      this.flushTimer = null;
    }
  }

  public dispose(): void {
    this.cleanup();
    this.statusBarItem.dispose();
  }

  public getTrackingStatus(): boolean {
    return this.tracking;
  }
}
