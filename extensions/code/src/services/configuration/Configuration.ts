import * as vscode from "vscode";

export interface Config {
  url: string;
  enabled: boolean;
  autoStart: boolean;
}

export class Configuration {
  private static instance: Configuration;
  private configChangeListener: vscode.Disposable | undefined = undefined;

  private constructor() {
    this.setupConfigListener();
  }

  public static getInstance(): Configuration {
    if (!Configuration.instance) {
      Configuration.instance = new Configuration();
    }
    return Configuration.instance;
  }

  public getConfig(): Config {
    const config = vscode.workspace.getConfiguration("flowtide");

    return {
      url: config.get<string>("url", "http://localhost:8080"),
      enabled: config.get<boolean>("enabled", true),
      autoStart: config.get<boolean>("autoStart", true),
    };
  }

  public async updateConfig(updates: Partial<Config>): Promise<void> {
    const config = vscode.workspace.getConfiguration("flowtide");

    for (const [key, value] of Object.entries(updates)) {
      if (key === "tracking" && typeof value === "object") {
        for (const [trackingKey, trackingValue] of Object.entries(value)) {
          await config.update(
            `tracking.${trackingKey}`,
            trackingValue,
            vscode.ConfigurationTarget.Global
          );
        }
      } else {
        await config.update(key, value, vscode.ConfigurationTarget.Global);
      }
    }
  }

  public async testConnection(): Promise<{ success: boolean; error?: string }> {
    const config = this.getConfig();

    try {
      const controller = new AbortController();
      const timeoutId = setTimeout(() => controller.abort(), 5000);

      const response = await fetch(`${config.url}/api/health`, {
        method: "GET",
        signal: controller.signal,
      });

      clearTimeout(timeoutId);

      if (response.ok) {
        return { success: true };
      } else {
        return {
          success: false,
          error: `Server responded with status ${response.status}`,
        };
      }
    } catch (error) {
      return {
        success: false,
        error: error instanceof Error ? error.message : "Connection failed",
      };
    }
  }

  private setupConfigListener(): void {
    this.configChangeListener = vscode.workspace.onDidChangeConfiguration(
      (event) => {
        if (event.affectsConfiguration("flowtide")) {
          this.onConfigChanged();
        }
      }
    );
  }

  private onConfigChanged(): void {
    vscode.commands.executeCommand("flowtide.revalidateConfiguration");
  }

  public dispose(): void {
    this.configChangeListener?.dispose();
  }
}
