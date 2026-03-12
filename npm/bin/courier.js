#!/usr/bin/env node

"use strict";

const path = require("path");
const { execFileSync } = require("child_process");

const binary = process.platform === "win32" ? "courier.exe" : "courier";
const binPath = path.join(__dirname, binary);

try {
  execFileSync(binPath, process.argv.slice(2), { stdio: "inherit" });
} catch (err) {
  if (err.status !== undefined) {
    process.exit(err.status);
  }
  console.error(`Failed to run Courier CLI: ${err.message}`);
  console.error(`Expected binary at: ${binPath}`);
  console.error(`Try reinstalling: npm install -g @trycourier/cli`);
  process.exit(1);
}
