import os from "os";
import fs from "fs";
import dotenv from "dotenv";

export default async (): Promise<boolean> => {
  const paths = [
    `${process.cwd()}/.courier`, // current working directory
    `${os.homedir()}/.courier` // user home directory
  ]

  for (const path of paths) {
    if (fs.existsSync(path)) {
      dotenv.config({ path });
      return true;
    }
  }
  return false;
}