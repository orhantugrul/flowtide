import * as vscode from "vscode";

export enum StatusBarState {
  ACTIVE,
  INACTIVE,
}

export class StatusBar implements vscode.Disposable {
  private _item: vscode.StatusBarItem;

  constructor(position = vscode.StatusBarAlignment.Right, priority = 100) {
    this._item = vscode.window.createStatusBarItem(position, priority);
    this._item.command = "flowtide.toggle";
    this._item.show();
    this.render(StatusBarState.INACTIVE);
  }

  /**
   * Renders the status bar item based on the current state
   * @param state The current state of the status bar
   */
  render(state: StatusBarState) {
    const config = {
      [StatusBarState.ACTIVE]: {
        text: "$(pulse) Flowtide",
        tooltip: "Click to stop tracking",
        backgroundColor: new vscode.ThemeColor(
          "statusBarItem.prominentBackground",
        ),
      },
      [StatusBarState.INACTIVE]: {
        text: "$(circle-outline) Flowtide",
        tooltip: "Click to start tracking",
        backgroundColor: undefined,
      },
    }[state];

    this._item.text = config.text;
    this._item.tooltip = config.tooltip;
    this._item.backgroundColor = config.backgroundColor;
  }

  /**
   * Disposes the status bar item and cleans up resources
   */
  dispose() {
    this._item.dispose();
  }
}
