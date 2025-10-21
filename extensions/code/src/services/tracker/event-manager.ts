import * as vscode from "vscode";

type OnTextEditorChange = (event?: vscode.TextEditor) => void;
type OnDocumentChange = (event: vscode.TextDocumentChangeEvent) => void;
type OnWindowStateChange = (event: vscode.WindowState) => void;

export class EventManager implements vscode.Disposable {
  private disposables: vscode.Disposable[] = [];

  onTextEditorChange(onChange: OnTextEditorChange): void {
    const disposable = vscode.window.onDidChangeActiveTextEditor(onChange);
    this.disposables.push(disposable);
  }

  onTextDocumentChange(onChange: OnDocumentChange): void {
    const disposable = vscode.workspace.onDidChangeTextDocument(onChange);
    this.disposables.push(disposable);
  }

  onWindowStateChange(onChange: OnWindowStateChange): void {
    const disposable = vscode.window.onDidChangeWindowState(onChange);
    this.disposables.push(disposable);
  }

  dispose(): void {
    this.disposables.forEach((disposable) => disposable.dispose());
    this.disposables = [];
  }
}
