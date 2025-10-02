import * as vscode from "vscode";
import { Configuration } from "./services/configuration/Configuration";
import { Logger } from "./services/logger/Logger";
import { Tracker } from "./services/tracker/Tracker";

export async function activate(context: vscode.ExtensionContext) {
  const configuration = Configuration.getInstance();
  const logger = Logger.getInstance();
  const tracker = Tracker.getInstance();

  logger.info("Flowtide starting");

  try {
    const connection = await configuration.testConnection();
    if (!connection.success) {
      logger.warn("Failed to connect to Flowtide server:", connection.error);
      vscode.window.showWarningMessage(
        `Cannot connect to Flowtide server: ${connection.error}. Please check your configuration.`
      );
    }

    registerCommands(context);

    const config = configuration.getConfig();
    if (config.enabled && config.autoStart) {
      await tracker.startTracking();
    }

    logger.info("Flowtide extension initialization completed");
  } catch (error) {
    logger.error("Failed to activate Flowtide extension:", error);
    vscode.window.showErrorMessage(
      "Failed to activate Flowtide extension. Check the output panel for details."
    );
  }
}

function registerCommands(context: vscode.ExtensionContext): void {
  const logger = Logger.getInstance();
  const tracker = Tracker.getInstance();

  const startCommand = vscode.commands.registerCommand(
    "flowtide.start",
    async () => {
      try {
        await tracker.startTracking();
      } catch (error) {
        logger.error("Failed to start tracking:", error);
        vscode.window.showErrorMessage("Failed to start tracking");
      }
    }
  );

  // Stop tracking command
  const stopCommand = vscode.commands.registerCommand(
    "flowtide.stop",
    async () => {
      try {
        await tracker.stopTracking();
      } catch (error) {
        logger.error("Failed to stop tracking:", error);
        vscode.window.showErrorMessage("Failed to stop tracking");
      }
    }
  );

  // Toggle tracking command
  const toggleCommand = vscode.commands.registerCommand(
    "flowtide.toggle",
    async () => {
      try {
        await Tracker.getInstance().toggleTracking();
      } catch (error) {
        logger.error("Failed to toggle tracking:", error);
        vscode.window.showErrorMessage("Failed to toggle tracking");
      }
    }
  );

  const dashboardCommand = vscode.commands.registerCommand(
    "flowtide.dashboard",
    async () => {
      try {
        const config = Configuration.getInstance().getConfig();

        if (!config.url) {
          vscode.window.showErrorMessage("Flowtide URL is not configured!");
          return;
        }

        await vscode.env.openExternal(vscode.Uri.parse(config.url));
        vscode.window.showInformationMessage("Opening Flowtide dashboard...");
      } catch (error) {
        logger.error("Failed to open dashboard:", error);
        vscode.window.showErrorMessage("Failed to open dashboard");
      }
    }
  );

  const logsCommand = vscode.commands.registerCommand("flowtide.logs", () => {
    logger.showOutputChannel();
  });

  const ravalideConfiguration = vscode.commands.registerCommand(
    "flowtide.revalidateConfiguration",
    async () => {
      try {
        if (Tracker.getInstance().getTrackingStatus()) {
          const connection = await Configuration.getInstance().testConnection();
          if (!connection.success) {
            logger.warn(
              "Connection lost after config change:",
              connection.error
            );
          }
        }
      } catch (error) {
        logger.error("Failed to handle config change:", error);
      }
    }
  );

  context.subscriptions.push(
    startCommand,
    stopCommand,
    toggleCommand,
    dashboardCommand,
    logsCommand,
    ravalideConfiguration
  );
}

export function deactivate() {
  Configuration.getInstance().dispose();
  Logger.getInstance().dispose();
  Tracker.getInstance().dispose();
}
