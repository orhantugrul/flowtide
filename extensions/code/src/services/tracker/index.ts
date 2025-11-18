import type * as vscode from "vscode";
import { getCurrentEditor, getCurrentProject } from "../../utils/workspace";
import { BufferManager } from "./buffer-manager";
import { EventManager } from "./event-manager";
import { StatusBar, StatusBarState } from "./status-bar";

export class Tracker implements vscode.Disposable {
  private static _instance: Tracker;

  /**
   * Gets the singleton instance of the tracker
   * @returns The global tracker instance
   */
  public static get instance(): Tracker {
    if (!Tracker._instance) {
      Tracker._instance = new Tracker();
    }
    return Tracker._instance;
  }

  private _bufferManager: BufferManager;
  private _eventManager: EventManager;
  private _statusBar: StatusBar;
  private _active = false;

  private constructor() {
    this._bufferManager = new BufferManager();
    this._eventManager = new EventManager();
    this._statusBar = new StatusBar();

    this._eventManager.onTextEditorChange(this.trackEditor);
    this._eventManager.onTextDocumentChange(this.trackDocument);
    this._eventManager.onWindowStateChange(this.trackWindow);
  }

  /**
   * Gets the buffer manager instance
   * @returns The buffer manager
   */
  public get bufferManager(): BufferManager {
    return this._bufferManager;
  }

  /**
   * Gets whether the tracker is currently active
   * @returns True if the tracker is active, false otherwise
   */
  public get active(): boolean {
    return this._active;
  }

  /**
   * Starts activity tracking with configuration validation
   * Initializes all tracking components and begins monitoring
   * @return Promise that resolves when tracking has started
   */
  public async start(): Promise<void> {
    if (this._active) {
      return;
    }

    await this._bufferManager.start();
    this._active = true;
    this._statusBar.render(StatusBarState.ACTIVE);
  }

  /**
   * Stops activity tracking and flushes remaining data
   * Gracefully shuts down all tracking components
   * @return Promise that resolves when tracking has stopped
   */
  public async stop(): Promise<void> {
    if (!this._active) {
      return;
    }

    await this._bufferManager.stop();
    this._active = false;
    this._statusBar.render(StatusBarState.INACTIVE);
  }

  /**
   * Tracks activity for a specific editor instance
   * @param editor The VS Code text editor to track
   * @return Promise that resolves when tracking is complete
   */
  private async trackEditor(editor?: vscode.TextEditor): Promise<void> {
    if (!editor || !this._active) {
      return;
    }

    const filePath = editor.document.uri.fsPath;
    const language = editor.document.languageId;
    const now = new Date().toISOString();

    const exists = this._bufferManager.buffer.items[filePath];
    if (exists) {
      this._bufferManager.buffer.update(filePath, { endTime: now });
      return;
    }

    try {
      const { id: projectId } = await getCurrentProject();
      const { id: editorId } = await getCurrentEditor();

      this._bufferManager.buffer.add({
        projectId,
        editorId,
        filePath,
        language,
        startTime: now,
        endTime: now,
      });
    } catch (error) {
      console.error("Failed to track editor:", error);
    }
  }

  /**
   * Tracks document changes and updates the buffer
   * @param event The document change event
   */
  private trackDocument(event: vscode.TextDocumentChangeEvent) {
    if (!this._active) {
      return;
    }

    const filePath = event.document.uri.fsPath;
    const now = new Date().toISOString();

    const exists = this._bufferManager.buffer.items[filePath];
    if (exists) {
      this._bufferManager.buffer.update(filePath, { endTime: now });
    }
  }

  /**
   * Tracks window state changes
   * @param event The window state change event
   */
  private trackWindow(event: vscode.WindowState) {
    if (!this._active) {
      return;
    }

    const state = {
      true: StatusBarState.ACTIVE,
      false: StatusBarState.INACTIVE,
    }[event.focused ? "true" : "false"];

    this._statusBar.render(state);

    const now = new Date().toISOString();
    for (const key of Object.keys(this._bufferManager.buffer.items)) {
      this._bufferManager.buffer.update(key, { endTime: now });
    }
  }

  /**
   * Disposes of all resources and stops tracking
   */
  public dispose() {
    this.stop();
    this._bufferManager.dispose();
    this._eventManager.dispose();
    this._statusBar.dispose();
  }
}
