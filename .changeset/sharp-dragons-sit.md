---
"@raulfdm/node-versions": minor
---

change from executable to node.

When compiling in a ubuntu environment (CI) and trying to run on macOS it throws an error:

```
cannot execute binary file
```

So, I'm temp. falling back to traditional .js + node execution.
