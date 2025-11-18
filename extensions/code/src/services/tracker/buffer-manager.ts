import type * as vscode from "vscode";
import { createActivities } from "../../api/activities";
import { Buffer } from "./buffer";

export class BufferManager implements vscode.Disposable {
  private _timer?: NodeJS.Timeout;
  private _buffer: Buffer = new Buffer();

  /**
   * Gets the current buffer as a read-only object
   * @returns The current buffer
   */
  get buffer(): Readonly<Buffer> {
    return this._buffer;
  }

  /**
   * Starts the buffer manager's periodic flush operation
   * @param period The flush period in milliseconds (default: 60000 ms)
   */
  async start(period = 60_000) {
    if (!this._timer) {
      this._timer = setInterval(() => this.flush(), period);
    }
  }

  /**
   * Stops the buffer manager and flushes remaining data
   */
  async stop() {
    if (this._timer) {
      clearInterval(this._timer);
      this._timer = undefined;
    }

    await this.flush();
  }

  /**
   * Flushes the buffered activities to the backend
   * @returns A promise that resolves when the flush operation is complete
   */
  async flush(): Promise<void> {
    const keys = Object.keys(this._buffer.items);
    if (keys.length === 0) {
      return;
    }

    try {
      const values = Object.values(this._buffer.items);
      const result = await createActivities(values);
      if (result.error) {
        throw Error(result.error.message);
      }

      this._buffer = new Buffer();
    } catch (error) {
      console.error("Buffer flush error:", error);
    }
  }

  /**
   * Disposes the buffer manager by stopping it
   */
  dispose() {
    this.stop();
  }
}
