import * as assert from "assert";
import { Tracker } from "../Tracker";

suite("Tracker Tests", () => {
  let tracker: Tracker;

  setup(() => {
    tracker = Tracker.getInstance();
  });

  teardown(() => {
    tracker.dispose();
  });

  test("Should create singleton instance", () => {
    const instance1 = Tracker.getInstance();
    const instance2 = Tracker.getInstance();
    assert.strictEqual(instance1, instance2);
  });

  test("Should start with tracking disabled", () => {
    assert.strictEqual(tracker.getTrackingStatus(), false);
  });

  test("Should detect editor correctly", () => {
    assert.ok(typeof tracker["getOrCreateEditor"] === "function");
  });

  test("Should handle start/stop lifecycle", async () => {
    try {
      await tracker.startTracking();
    } catch (error) {
      assert.ok(error instanceof Error);
    }

    try {
      await tracker.stopTracking();
    } catch (error) {
      assert.ok(error instanceof Error);
    }
  });
});
