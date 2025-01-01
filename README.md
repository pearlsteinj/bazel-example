# bazel-example

![build](https://github.com/pearlsteinj/bazel-example/actions/workflows/build.yaml/badge.svg)

## Overview

Contains a minimal example of setting up bazel using `MODULE.bazel` files, used as a base for various projects. 

Includes: 

- Set up of a bazel project that mixes `golang` / `proto` / `nodejs`. 
- Github actions to act as minimal CI.
- Examples of building an AWS lambda and grpc backends in go.
- Examples of building oci container images in `go` and `nodejs`.

## Commands

To update go dependencies when adding additional entries to the `go.mod` file, run: 
```
bazel run @rules_go//go -- mod tidy
```

To update go dependencies when adding additional npm entries to the `package.json` file, run:
```
bazel run -- @pnpm --dir $PWD install
```