#!/usr/bin/env node

import * as path from "node:path";
import * as fs from "node:fs";
import https from "https";
import tar from "tar";
import { pipeline } from "stream/promises";

// Mapping from Node's `process.arch` to Golang's `$GOARCH`
const ARCH_MAPPING = {
	ia32: "386",
	x64: "amd64",
	arm: "arm",
	arm64: "arm64",
};

// Mapping between Node's `process.platform` to Golang's
const PLATFORM_MAPPING = {
	darwin: "darwin",
	linux: "linux",
	win32: "windows",
	freebsd: "freebsd",
};

const command = process.argv[2];

if (command === "install") {
	await install();
} else if (command === "uninstall") {
	// do something
} else {
	console.log(
		"Invalid command. 'install' and 'uninstall' are the only supported commands",
	);
	process.exit(1);
}

async function install() {
	console.log("Installing binary...");
	validateOsAndArch();
	const pkgJson = readPackageJson();
	const metaData = getMetaData(pkgJson);
	createBinPath(metaData.binPath);

	try {
		await downloadFile(metaData.url, metaData.binTarGz);
		await tar.x({
			file: metaData.binTarGz,
			cwd: metaData.binPath,
		});

		console.log("Binary installed successfully");
	} catch (error) {
		console.error(`Error downloading binary: ${error}`);
		process.exit(1);
	}
}

function readPackageJson() {
	const packageJsonPath = path.join(".", "package.json");

	if (!fs.existsSync(packageJsonPath)) {
		console.error("package.json not found in current directory");
		process.exit(1);
	}

	const packageJson = JSON.parse(fs.readFileSync(packageJsonPath, "utf-8"));
	const err = validateConfiguration(packageJson);

	if (err) {
		console.error(err);
		process.exit(1);
	}

	return packageJson;
}

/**
 * @typedef {Object} GoBinary
 * @property {(string|undefined)} path - The path of the Go binary.
 * @property {(string|undefined)} name - The name of the Go binary.
 * @property {(string|undefined)} url - The URL of the Go binary.
 */

/**
 * @typedef {Object} PackageJson
 * @property {(GoBinary|undefined)} goBinary - The goBinary object.
 * @property {string} version - The version of the package.
 */

/**
 * @typedef {Object} MetaData
 * @property {string} binName - The name of the binary.
 * @property {string} binPath - The path of the binary.
 * @property {string} url - The URL of the binary.
 * @property {string} version - The version of the binary.
 */

/**
 * Extracts metadata from a package.json object.
 *
 * @param {PackageJson} packageJson - The package.json object.
 *
 * @returns {MetaData} An object containing the binary name, path, URL, and version.
 */
function getMetaData(packageJson) {
	const binPath = packageJson.goBinary.path;
	let binName = packageJson.goBinary.name;
	let url = packageJson.goBinary.url;
	let version = packageJson.version;

	if (version[0] === "v") {
		version = version.substring(1); // strip the 'v' if necessary v0.0.1 => 0.0.1
	}

	// Binary name on Windows has .exe suffix
	if (process.platform === "win32") {
		binName += ".exe";
	}

	// Interpolate variables in URL, if necessary
	url = url.replace(/{{arch}}/g, ARCH_MAPPING[process.arch]);
	url = url.replace(/{{platform}}/g, PLATFORM_MAPPING[process.platform]);
	url = url.replace(/{{version}}/g, version);
	url = url.replace(/{{bin_name}}/g, binName);

	return {
		binName,
		binPath,
		binFullName: path.join(process.cwd(), binPath),
		get binTarGz() {
			return `${this.binFullName}.tar.gz`;
		},
		url,
		version,
	};
}

function validateOsAndArch() {
	if (!(process.arch in ARCH_MAPPING)) {
		console.error(`Invalid architecture: ${process.arch}`);
		process.exit(1);
	}

	if (!(process.platform in PLATFORM_MAPPING)) {
		console.error(`Invalid platform: ${process.platform} `);
		process.exit(1);
	}
}

/**
 * Validates the package.json object.
 * @param {PackageJson} packageJson - The package.json object.
 * @returns {string} An error message if the package.json object is invalid.
 */
function validateConfiguration(packageJson) {
	if (!packageJson.version) {
		return "'version' property must be specified";
	}

	if (!packageJson.goBinary || typeof packageJson.goBinary !== "object") {
		return "'goBinary' property must be defined and be an object";
	}

	if (!packageJson.goBinary.name) {
		return "'name' property is necessary";
	}

	if (!packageJson.goBinary.path) {
		return "'path' property is necessary";
	}

	if (!packageJson.goBinary.url) {
		return "'url' property is required";
	}
}

/**
 * Creates a directory at the specified path.
 * @param {string} binPath - The path of the directory to create.
 */
function createBinPath(binPath) {
	if (!fs.existsSync(binPath)) {
		fs.mkdirSync(binPath, { recursive: true });
	}
}

/**
 * Downloads a file from a given URL and saves it to the specified path.
 *
 * @param {string} url - The URL of the file to download.
 * @param {string} outputPath - The path where the downloaded file should be saved.
 * @returns {Promise<string>} A promise that resolves with the path of the downloaded file.
 * @throws {Error} Throws an error if the download or file writing fails.
 */
async function downloadFile(url, outputPath) {
	return new Promise((resolve, reject) => {
		const processResponse = (response) => {
			// Check if the response is a redirect
			if (
				response.statusCode >= 300 &&
				response.statusCode < 400 &&
				response.headers.location
			) {
				https
					.get(response.headers.location, processResponse)
					.on("error", reject);
			} else if (response.statusCode === 200) {
				const fileStream = fs.createWriteStream(outputPath);
				pipeline(response, fileStream)
					.then(() => resolve(outputPath))
					.catch((error) => {
						reject(`Error during download: ${error.message}`);
					});
			} else {
				reject(
					`Server responded with ${response.statusCode}: ${response.statusMessage}`,
				);
			}
		};

		https.get(url, processResponse).on("error", reject);
	});
}
