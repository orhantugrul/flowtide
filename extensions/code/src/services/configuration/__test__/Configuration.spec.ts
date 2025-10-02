import * as assert from "assert";
import { Configuration } from "../Configuration";

suite("Configuration Tests", () => {
  let configService: Configuration;

  setup(() => {
    configService = Configuration.getInstance();
  });

  teardown(() => {
    configService.dispose();
  });

  test("Should create singleton instance", () => {
    const instance1 = Configuration.getInstance();
    const instance2 = Configuration.getInstance();
    assert.strictEqual(instance1, instance2);
  });

  test("Should get default configuration", () => {
    const config = configService.getConfig();

    assert.ok(config.url);
    assert.strictEqual(typeof config.enabled, "boolean");
    assert.strictEqual(typeof config.autoStart, "boolean");
  });
});
