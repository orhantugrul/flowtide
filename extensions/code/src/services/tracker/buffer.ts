import type { Activity } from "../../types/activity";

type Key = string;
type Value = Omit<Activity, "id" | "createdAt" | "updatedAt" | "deletedAt">;

export class Buffer {
  private _items: Record<Key, Value> = {};

  /**
   * Gets the buffered items as a read-only record
   * @returns The buffered items
   */
  get items(): Record<string, Readonly<Value>> {
    return this._items;
  }

  /**
   * Adds a new item to the buffer
   * @param value The item values to add
   */
  add(value: Value) {
    this._items[value.filePath] = {
      projectId: value.projectId,
      editorId: value.editorId,
      filePath: value.filePath,
      language: value.language,
      startTime: value.startTime,
      endTime: value.endTime,
    };
  }

  /**
   * Updates an existing item in the buffer
   * @param key The key of the item to update
   * @param value Partial values to update
   */
  update(key: string, value: Partial<Value>) {
    if (this._items[key]) {
      this._items[key] = { ...this._items[key], ...value };
    }
  }
}
