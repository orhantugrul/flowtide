import * as vscode from "vscode";

export enum LogLevel {
  DEBUG = 0,
  INFO = 1,
  WARN = 2,
  ERROR = 3,
}

export class Logger {
  private static instance: Logger;
  private outputChannel: vscode.OutputChannel;
  private logLevel: LogLevel = LogLevel.INFO;

  private constructor() {
    this.outputChannel = vscode.window.createOutputChannel("Flowtide");
    this.setupLogLevel();
  }

  public static getInstance(): Logger {
    if (!Logger.instance) {
      Logger.instance = new Logger();
    }
    return Logger.instance;
  }

  private setupLogLevel(): void {
    const config = vscode.workspace.getConfiguration("flowtide");
    const level = config.get<string>("logLevel", "info").toLowerCase();

    switch (level) {
      case "debug":
        this.logLevel = LogLevel.DEBUG;
        break;
      case "info":
        this.logLevel = LogLevel.INFO;
        break;
      case "warn":
        this.logLevel = LogLevel.WARN;
        break;
      case "error":
        this.logLevel = LogLevel.ERROR;
        break;
      default:
        this.logLevel = LogLevel.INFO;
    }
  }

  private log(level: LogLevel, message: string, ...args: any[]): void {
    if (level < this.logLevel) {
      return;
    }

    const timestamp = new Date().toISOString();
    const levelName = LogLevel[level];
    const formattedMessage = `[${timestamp}] [${levelName}] ${message}`;

    this.outputChannel.appendLine(formattedMessage);

    if (args.length > 0) {
      this.outputChannel.appendLine(JSON.stringify(args, null, 2));
    }

    // Also log to console in development
    if (process.env.NODE_ENV === "development") {
      console.log(formattedMessage, ...args);
    }
  }

  public debug(message: string, ...args: any[]): void {
    this.log(LogLevel.DEBUG, message, ...args);
  }

  public info(message: string, ...args: any[]): void {
    this.log(LogLevel.INFO, message, ...args);
  }

  public warn(message: string, ...args: any[]): void {
    this.log(LogLevel.WARN, message, ...args);
  }

  public error(message: string, ...args: any[]): void {
    this.log(LogLevel.ERROR, message, ...args);
  }

  public showOutputChannel(): void {
    this.outputChannel.show();
  }

  public dispose(): void {
    this.outputChannel.dispose();
  }
}
