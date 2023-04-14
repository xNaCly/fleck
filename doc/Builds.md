# Builds

Fleck comes in two builds, the default fully featured build and the `bare` build.

To get the bare build, build the binary using:

```bash
go build -tags=bare .
```

## Differences

The `bare` build does not support:

- cli options and flags
- colored output
- extensive logs and time stamps for compilation steps
- templates
- styling
- preprocessor

The `bare` builds reason for existence is to be light weight and only include the main feature of fleck, the markdown to html conversion.
