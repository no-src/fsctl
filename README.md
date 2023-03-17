# fsctl

[![Build](https://img.shields.io/github/actions/workflow/status/no-src/fsctl/go.yml?branch=main)](https://github.com/no-src/fsctl/actions)
[![License](https://img.shields.io/github/license/no-src/fsctl)](https://github.com/no-src/fsctl/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/no-src/fsctl.svg)](https://pkg.go.dev/github.com/no-src/fsctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/no-src/fsctl)](https://goreportcard.com/report/github.com/no-src/fsctl)
[![codecov](https://codecov.io/gh/no-src/fsctl/branch/main/graph/badge.svg?token=BTPKR8G6QI)](https://codecov.io/gh/no-src/fsctl)
[![Release](https://img.shields.io/github/v/release/no-src/fsctl)](https://github.com/no-src/fsctl/releases)

English | [简体中文](README-CN.md)

The fsctl is a configuration-based file operation and validation tool.

## Installation

The first need [Go](https://go.dev/doc/install) installed (**version 1.19+ is required**), then you can use the below
command to install `fsctl`.

```bash
go install github.com/no-src/fsctl/...@latest
```

### Run In Docker

You can use the [build-docker.sh](/scripts/build-docker.sh) script to build the docker image and you should clone this
repository and `cd` to the root path of the repository first.

```bash
$ ./scripts/build-docker.sh
```

Or pull the docker image directly from [DockerHub](https://hub.docker.com/r/nosrc/fsctl) with the command below.

```bash
$ docker pull nosrc/fsctl
```

For more scripts about release and docker, see the [scripts](/scripts) directory.

## Quick Start

```bash
$ fsctl -conf fsctl.yaml
```

## For More Information

### Help Info

```bash
$ fsctl -h
```

### Version Info

```bash
$ fsctl -v
```

### About Info

```bash
$ fsctl -about
```