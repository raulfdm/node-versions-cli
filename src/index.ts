import meow from "meow";
import consola from "consola";

import { NodeVersions } from "./schema";

const response = await fetch("https://nodejs.org/dist/index.json");

const nodeVersionsJson = await response.json();

const nodeVersions = NodeVersions.parse(nodeVersionsJson);

const { flags } = meow(
	`Usage
$ node-versions <flags>

Options
--lts   Show all LTS versions
--all   Show all versions

Examples
$ node-versions
v21.1

$ node-versions --lts
v20.11.0
v20.10.0
v20.9.0
v18.19.0
v18.18.2
...

$ node-versions --all
v21.6.1
v21.6.0
v21.5.0
v21.4.0
v21.3.0
v21.2.0
v21.1.0
v21.0.0
v20.11.0
...
`,
	{
		importMeta: import.meta,
		flags: {
			lts: {
				type: "boolean",
				default: false,
			},
			all: {
				type: "boolean",
				default: false,
			},
		},
	},
);

if (flags.lts) {
	showLts();
} else if (flags.all) {
	showAll();
} else {
	const [latestVersions] = nodeVersions;
	logVersions([latestVersions]);
}

function showAll() {
	consola.log("All Versions:");
	logVersions(nodeVersions);
}

function showLts() {
	consola.info("LTS Versions:");
	const ltsVersions = nodeVersions.filter((version) => version.lts);
	logVersions(ltsVersions);
}

function logVersions(nodeVersions: NodeVersions) {
	const result = nodeVersions.reduce((acc, nodeVersion) => {
		return `${acc}\n${nodeVersion.version}`;
	}, "");

	consola.log(result);
}
