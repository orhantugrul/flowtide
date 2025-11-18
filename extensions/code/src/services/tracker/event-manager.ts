import * as vscode from "vscode";

type OnTextEditorChange = (event?: vscode.TextEditor) => void;
type OnDocumentChange = (event: vscode.TextDocumentChangeEvent) => void;
type OnWindowStateChange = (event: vscode.WindowState) => void;

export class EventManager implements vscode.Disposable {
  private _disposables: vscode.Disposable[] = [];

  /**
   * Registers a callback to be invoked when the active text editor changes
   * @param onChange The callback to invoke when the active text editor changes
   */
  onTextEditorChange(onChange: OnTextEditorChange): void {
    const disposable = vscode.window.onDidChangeActiveTextEditor(onChange);
    this._disposables.push(disposable);
  }

  /**
   * Registers a callback to be invoked when a text document changes
   * @param onChange The callback to invoke when a text document changes
   */
  onTextDocumentChange(onChange: OnDocumentChange): void {
    const disposable = vscode.workspace.onDidChangeTextDocument(onChange);
    this._disposables.push(disposable);
  }

  /**
   * Registers a callback to be invoked when the window state changes
   * @param onChange The callback to invoke when the window state changes
   */
  onWindowStateChange(onChange: OnWindowStateChange): void {
    const disposable = vscode.window.onDidChangeWindowState(onChange);
    this._disposables.push(disposable);
  }

  /**
   * Disposes all registered event listeners
   * Cleans up resources used by the EventManager
   */
  dispose(): void {
    this._disposables.forEach((disposable) => disposable.dispose());
    this._disposables = [];
  }
}
