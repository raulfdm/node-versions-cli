# @raulfdm/node-versions

## 1.1.0

### Minor Changes

- 3e374dd: change from executable to node.

  When compiling in a ubuntu environment (CI) and trying to run on macOS it throws an error:

  ```
  cannot execute binary file
  ```

  So, I'm temp. falling back to traditional .js + node execution.

## 1.0.1

### Patch Changes

- f2a893b: minify binary and publish only dist files

## 1.0.0

### Major Changes

- 43a7eea: Release CLI
