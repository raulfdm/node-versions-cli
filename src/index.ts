import { parseArgs } from "util";
import { NodeVersions } from "./schema";

const response = await fetch("https://nodejs.org/dist/index.json");

const nodeVersionsJson = await response.json();

const nodeVersions = NodeVersions.parse(nodeVersionsJson);

const { values: argValues } = parseArgs({
	args: Bun.argv,
	options: {
		lts: {
			type: "boolean",
			default: false,
		},
	},
	strict: true,
	allowPositionals: true,
});

if (argValues.lts) {
	showLts();
}

function showLts() {
	console.log("LTS Versions:");
	const ltsVersions = nodeVersions.filter((version) => version.lts);
	logVersions(ltsVersions);
}

function logVersions(nodeVersions: NodeVersions) {
	for (const nodeVersion of nodeVersions) {
		console.log(nodeVersion.version);
	}
}
