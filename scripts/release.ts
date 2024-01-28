import { $ } from "bun";
import { execaCommand } from "execa";
import consola from "consola";

consola.info("Starting release process");

try {
	consola.log("Trying to publish");
	await execaCommand("pnpm publish --no-git-checks", {
		shell: true,
		all: true,
	});
	consola.success("Published successfully");

	consola.log("Trying to push tags...");
	await $`git push --follow-tags`;
	consola.success("tags pushed successfully");
} catch (error) {
	if (error instanceof Error) {
		if (
			error.message.includes(
				"You cannot publish over the previously published versions",
			)
		) {
			console.info("Version already published, skipping...");
			process.exit(0);
		} else {
			consola.error("Something went wrong", error.message);
			process.exit(1);
		}
	} else {
		consola.error("Unknown error", error);
		process.exit(1);
	}
}

// const a = await $`git`.text();

// console.log(a);
