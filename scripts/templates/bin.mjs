#!/usr/bin/env node

import { execSync } from "child_process";
import { fileURLToPath } from "url";
import path from "path";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const fullPath = path.join(__dirname, "bin/node-versions-cli");

// Get the arguments passed to the script, excluding the first two elements
// The first element is the path to the node executable
// The second element is the path to your script
const args = process.argv.slice(2);

// Join the arguments into a single string
const argsString = args.join(" ");

execSync(`${fullPath} ${argsString}`, { stdio: "inherit" });
