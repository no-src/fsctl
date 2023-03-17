# fsctl

[![Build](https://img.shields.io/github/actions/workflow/status/no-src/fsctl/go.yml?branch=main)](https://github.com/no-src/fsctl/actions)
[![License](https://img.shields.io/github/license/no-src/fsctl)](https://github.com/no-src/fsctl/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/no-src/fsctl.svg)](https://pkg.go.dev/github.com/no-src/fsctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/no-src/fsctl)](https://goreportcard.com/report/github.com/no-src/fsctl)
[![codecov](https://codecov.io/gh/no-src/fsctl/branch/main/graph/badge.svg?token=BTPKR8G6QI)](https://codecov.io/gh/no-src/fsctl)
[![Release](https://img.shields.io/github/v/release/no-src/fsctl)](https://github.com/no-src/fsctl/releases)

[English](README.md) | 简体中文

fsctl是一种基于配置的文件操作和验证工具

## 安装

首先需要确保已经安装了[Go](https://golang.google.cn/doc/install) (**版本必须是1.19+**)，
然后你就可以使用下面的命令来安装`fsctl`了

```bash
go install github.com/no-src/fsctl/...@latest
```

### 在Docker中运行

你可以使用[build-docker.sh](/scripts/build-docker.sh)脚本来构建docker镜像，首先你需要克隆本仓库并且`cd`到本仓库的根目录

```bash
$ ./scripts/build-docker.sh
```

或者使用以下命令直接从[DockerHub](https://hub.docker.com/r/nosrc/fsctl)中拉取docker镜像

```bash
$ docker pull nosrc/fsctl
```

更多关于发布与docker的脚本参见[scripts](scripts)目录

## 快速开始

```bash
$ fsctl -conf fsctl.yaml
```

## 更多信息

### 帮助信息

```bash
$ fsctl -h
```

### 版本信息

```bash
$ fsctl -v
```

### 关于信息

```bash
$ fsctl -about
```