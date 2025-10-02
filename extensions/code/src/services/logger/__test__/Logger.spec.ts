import * as assert from "assert";
import { Logger, LogLevel } from "../Logger";

suite("Logger Tests", () => {
  let logger: Logger;

  setup(() => {
    logger = Logger.getInstance();
  });

  teardown(() => {
    logger.dispose();
  });

  test("Should create singleton instance", () => {
    const instance1 = Logger.getInstance();
    const instance2 = Logger.getInstance();
    assert.strictEqual(instance1, instance2);
  });

  test("Should have correct log levels", () => {
    assert.strictEqual(LogLevel.DEBUG, 0);
    assert.strictEqual(LogLevel.INFO, 1);
    assert.strictEqual(LogLevel.WARN, 2);
    assert.strictEqual(LogLevel.ERROR, 3);
  });

  test("Should have logging methods", () => {
    assert.ok(typeof logger.debug === "function");
    assert.ok(typeof logger.info === "function");
    assert.ok(typeof logger.warn === "function");
    assert.ok(typeof logger.error === "function");
    assert.ok(typeof logger.showOutputChannel === "function");
  });

  test("Should handle different log levels", () => {
    assert.doesNotThrow(() => logger.debug("Debug message"));
    assert.doesNotThrow(() => logger.info("Info message"));
    assert.doesNotThrow(() => logger.warn("Warn message"));
    assert.doesNotThrow(() => logger.error("Error message"));
  });
});
