#!/usr/bin/env bash

# update repository
#git pull --no-rebase

# update the last git commit
echo -e "$(git rev-parse main)\c" >internal/version/commit

# set GOPROXY environment variable
# export GOPROXY=https://goproxy.cn

############################## linux-release ##############################

# set go env for linux
export GOOS=linux
export GOARCH=amd64

# build fsctl
go build -v -o . ./...

export FSCTL_RELEASE_GO_VERSION=$(go version | awk '{print $3}')
export FSCTL_RELEASE_VERSION=$(./fsctl -v | awk 'NR==1 {print $3}')

# release path, for example, fsctl_go1.20.2_amd64_linux_v0.0.1
export FSCTL_RELEASE="fsctl_${FSCTL_RELEASE_GO_VERSION}_${GOARCH}_${GOOS}_${FSCTL_RELEASE_VERSION}"

rm -rf "$FSCTL_RELEASE"
mkdir "$FSCTL_RELEASE"
mv fsctl "$FSCTL_RELEASE/"

# linux release archive
tar -zcvf "$FSCTL_RELEASE.tar.gz" "$FSCTL_RELEASE"

rm -rf "$FSCTL_RELEASE"

############################## linux-release ##############################

############################# windows-release #############################

# set go env for windows
export GOOS=windows
export GOARCH=amd64

# build fsctl
go build -v -o . ./...

# release path, for example, fsctl_go1.20.2_amd64_windows_v0.0.1
export FSCTL_RELEASE="fsctl_${FSCTL_RELEASE_GO_VERSION}_${GOARCH}_${GOOS}_${FSCTL_RELEASE_VERSION}"

mkdir "$FSCTL_RELEASE"
mv fsctl.exe "$FSCTL_RELEASE/"

# windows release archive
zip -r "$FSCTL_RELEASE.zip" "$FSCTL_RELEASE"

rm -rf "$FSCTL_RELEASE"

############################# windows-release #############################

############################## macOS-release ##############################

# set go env for macOS
export GOOS=darwin
export GOARCH=amd64

# build fsctl
go build -v -o . ./...

# release path, for example, fsctl_go1.20.2_amd64_darwin_v0.0.1
export FSCTL_RELEASE="fsctl_${FSCTL_RELEASE_GO_VERSION}_${GOARCH}_${GOOS}_${FSCTL_RELEASE_VERSION}"

rm -rf "$FSCTL_RELEASE"
mkdir "$FSCTL_RELEASE"
mv fsctl "$FSCTL_RELEASE/"

# macOS release archive
tar -zcvf "$FSCTL_RELEASE.tar.gz" "$FSCTL_RELEASE"

rm -rf "$FSCTL_RELEASE"

############################## macOS-release ##############################

# reset commit file
echo -e "\c" >internal/version/commit
