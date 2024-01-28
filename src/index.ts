import meow from "meow";
import consola from "consola";
import groupBy from "just-group-by";
import { semver } from "bun";

import { NodeVersions } from "./schema";

const { flags, showHelp } = meow(
	`🌟 Node Versions CLI 🌟

Usage:
$ node-versions <flag>

Options:
--all		Show all versions
--all-lts		Show all LTS versions
--latest	Show latest version
--latest-of	Show latest version of a specific version
--lts		Show current LTS version

Examples:
$ node-versions --all
$ node-versions --all-lts
$ node-versions --latest
$ node-versions --latest-of 20
$ node-versions --lts
`,
	{
		importMeta: import.meta,
		flags: {
			lts: {
				type: "boolean",
			},
			allLts: {
				type: "boolean",
			},
			latest: {
				type: "boolean",
			},
			latestOf: {
				type: "string",
			},
		},
	},
);

const nodeVersions = await getNodeVersions();

if (flags.lts) {
	showLts();
} else if (flags.all) {
	showAll();
} else if (flags.allLts) {
	showAllLts();
} else if (flags.latestOf) {
	showLatestOf();
} else if (flags.latest) {
	showLatest();
} else {
	showHelp();
}

function showAll() {
	consola.log("All Versions:");
	logVersions(nodeVersions);
}

function showLts() {
	consola.info("Current LTS:");
	const [currentLTS] = nodeVersions.filter((version) => version.lts);
	logVersions([currentLTS]);
}

function showAllLts() {
	consola.info("LTS Versions:");
	const ltsVersions = nodeVersions.filter((version) => version.lts);
	logVersions(ltsVersions);
}

function logVersions(nodeVersions: NodeVersions) {
	const ascendingVersions = nodeVersions.sort((a, b) =>
		semver.order(a.version, b.version),
	);

	const result = ascendingVersions.reduce((acc, nodeVersion) => {
		return `${acc}${nodeVersion.version}\n`;
	}, "");

	consola.log(result.trim());
}

function showLatestOf() {
	const { latestOf } = flags;
	const prependVersion = `v${latestOf}`;

	const groupedVersions = groupBy(
		nodeVersions,
		(version) => version.version.split(".")[0],
	);

	const allVersionsOf = groupedVersions[prependVersion];

	if (!allVersionsOf) {
		consola.error(`No versions found for ${prependVersion}`);
		return;
	}

	const [latestVersion] = allVersionsOf;

	consola.info(`Latest version of ${prependVersion}:`);
	logVersions([latestVersion]);
}

function showLatest() {
	const [latestVersions] = nodeVersions;
	console.log("Latest version:");
	logVersions([latestVersions]);
}

async function getNodeVersions() {
	const response = await fetch("https://nodejs.org/dist/index.json");

	const nodeVersionsJson = await response.json();

	return NodeVersions.parse(nodeVersionsJson);
}
