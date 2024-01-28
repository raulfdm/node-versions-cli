import { NodeVersions } from "./schema";

const response = await fetch("https://nodejs.org/dist/index.json");

const nodeVersionsJson = await response.json();

const nodeVersions = NodeVersions.parse(nodeVersionsJson);

console.log(nodeVersions);
