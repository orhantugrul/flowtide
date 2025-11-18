import * as vscode from "vscode";
import { Tracker } from "./services/tracker";

export async function activate(context: vscode.ExtensionContext) {
  const tracker = Tracker.instance;
  const config = vscode.workspace.getConfiguration("flowtide");

  try {
    registerCommands(context);

    if (config.get<boolean>("autoStart")) {
      await tracker.start();
      vscode.window.showInformationMessage("Flowtide started traking");
    }
  } catch (error) {
    const message = `Failed to activate Flowtide extension: ${error}`;
    vscode.window.showErrorMessage(message);
  }
}

function registerCommands(context: vscode.ExtensionContext): void {
  const tracker = Tracker.instance;
  const config = vscode.workspace.getConfiguration("flowtide");

  const startCommand = vscode.commands.registerCommand(
    "flowtide.start",
    async () => {
      try {
        await tracker.start();
        vscode.window.showInformationMessage("Flowtide started tracking");
      } catch (error) {
        vscode.window.showErrorMessage("Failed to start tracking");
      }
    },
  );

  const stopCommand = vscode.commands.registerCommand(
    "flowtide.stop",
    async () => {
      try {
        await tracker.stop();
        vscode.window.showInformationMessage("Flowtide stopped tracking");
      } catch (error) {
        vscode.window.showErrorMessage("Failed to stop tracking");
      }
    },
  );

  const dashboardCommand = vscode.commands.registerCommand(
    "flowtide.dashboard",
    async () => {
      try {
        const url = config.get<string>("url", "http://localhost:8080");
        await vscode.env.openExternal(vscode.Uri.parse(url));
        vscode.window.showInformationMessage("Opening Flowtide dashboard...");
      } catch (error) {
        console.error("Failed to open dashboard:", error);
        vscode.window.showErrorMessage("Failed to open dashboard");
      }
    },
  );

  context.subscriptions.push(startCommand, stopCommand, dashboardCommand);
}

export function deactivate() {
  Tracker.instance.dispose();
}
