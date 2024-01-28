# node-versions CLI

A simple CLI to easily check node versions

## Getting started

Install it globally:

```bash
npm add -g @raulfdm/node-versions
```

Then run:

```bash
node-versions <flag>
```

## API

```
🌟 Node Versions CLI 🌟

Usage:
$ node-versions <flag>

Options:
--all         Show all versions
--all-lts     Show all LTS versions
--latest      Show latest version
--latest-of   Show latest version of a specific version
--lts         Show current LTS version

Examples:
$ node-versions --all
$ node-versions --all-lts
$ node-versions --latest
$ node-versions --latest-of 20
$ node-versions --lts
```

## Contributing

1. Make sure to have bun v1.0.25 or higher
2. clone this repository
3. install the dependencies:
   ```
   bun install
   ```
4. Do the changes on the files in `src/*`;
5. run the dev command:
   ```
   bun run dev
   ```

If your change must be published:

1. Run `bun changeset`
1. Select the type of version the change includes following semver logic
1. Commit the changeset file alongside with your changes
1. Open a PR
1. Wait until it gets reviewed and approved
