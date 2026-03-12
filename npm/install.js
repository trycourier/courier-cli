#!/usr/bin/env node

"use strict";

const https = require("https");
const http = require("http");
const fs = require("fs");
const path = require("path");
const { execSync } = require("child_process");

const REPO = "trycourier/courier-cli";
const VERSION = require("./package.json").version;
const BIN_DIR = path.join(__dirname, "bin");

const PLATFORM_MAP = {
  "darwin-arm64": { archive: `courier_${VERSION}_macos_arm64.zip`, binary: "courier" },
  "darwin-x64": { archive: `courier_${VERSION}_macos_amd64.zip`, binary: "courier" },
  "linux-x64": { archive: `courier_${VERSION}_linux_amd64.tar.gz`, binary: "courier" },
  "linux-arm64": { archive: `courier_${VERSION}_linux_arm64.tar.gz`, binary: "courier" },
  "linux-ia32": { archive: `courier_${VERSION}_linux_386.tar.gz`, binary: "courier" },
  "linux-arm": { archive: `courier_${VERSION}_linux_armv6.tar.gz`, binary: "courier" },
  "win32-x64": { archive: `courier_${VERSION}_windows_amd64.zip`, binary: "courier.exe" },
  "win32-arm64": { archive: `courier_${VERSION}_windows_arm64.zip`, binary: "courier.exe" },
  "win32-ia32": { archive: `courier_${VERSION}_windows_386.zip`, binary: "courier.exe" },
};

const key = `${process.platform}-${process.arch}`;
const target = PLATFORM_MAP[key];

if (!target) {
  console.error(
    `Unsupported platform: ${key}\n` +
    `Supported: ${Object.keys(PLATFORM_MAP).join(", ")}\n` +
    `You can download the binary directly from:\n` +
    `  https://github.com/${REPO}/releases/tag/v${VERSION}`
  );
  process.exit(1);
}

const url = `https://github.com/${REPO}/releases/download/v${VERSION}/${target.archive}`;
const archivePath = path.join(BIN_DIR, target.archive);

function follow(requestUrl, redirects) {
  if (redirects > 10) {
    throw new Error("Too many redirects");
  }
  return new Promise((resolve, reject) => {
    const lib = requestUrl.startsWith("https") ? https : http;
    lib.get(requestUrl, { headers: { "User-Agent": `@trycourier/cli-npm/${VERSION}` } }, (res) => {
      if (res.statusCode >= 300 && res.statusCode < 400 && res.headers.location) {
        resolve(follow(res.headers.location, redirects + 1));
        return;
      }
      if (res.statusCode !== 200) {
        reject(new Error(`Download failed: HTTP ${res.statusCode} from ${requestUrl}`));
        return;
      }
      resolve(res);
    }).on("error", reject);
  });
}

async function download() {
  console.log(`Downloading Courier CLI v${VERSION} for ${key}...`);

  const res = await follow(url, 0);
  const file = fs.createWriteStream(archivePath);
  await new Promise((resolve, reject) => {
    res.pipe(file);
    file.on("finish", () => { file.close(resolve); });
    file.on("error", reject);
  });
}

function extract() {
  const binPath = path.join(BIN_DIR, target.binary);

  if (target.archive.endsWith(".tar.gz")) {
    execSync(`tar -xzf "${archivePath}" -C "${BIN_DIR}" "${target.binary}"`, { stdio: "pipe" });
  } else if (target.archive.endsWith(".zip")) {
    if (process.platform === "win32") {
      execSync(
        `powershell -Command "Expand-Archive -Force '${archivePath}' '${BIN_DIR}'"`,
        { stdio: "pipe" }
      );
    } else {
      execSync(`unzip -o -j "${archivePath}" "${target.binary}" -d "${BIN_DIR}"`, { stdio: "pipe" });
    }
  }

  if (process.platform !== "win32") {
    fs.chmodSync(binPath, 0o755);
  }

  fs.unlinkSync(archivePath);

  console.log(`Courier CLI v${VERSION} installed successfully.`);
}

async function main() {
  try {
    fs.mkdirSync(BIN_DIR, { recursive: true });
    await download();
    extract();
  } catch (err) {
    console.error(
      `Failed to install Courier CLI: ${err.message}\n` +
      `You can download the binary manually from:\n` +
      `  https://github.com/${REPO}/releases/tag/v${VERSION}`
    );
    process.exit(1);
  }
}

main();
