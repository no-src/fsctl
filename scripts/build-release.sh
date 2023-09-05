#!/usr/bin/env bash

# update repository
#git pull --no-rebase

# update the last git commit
echo -e "$(git rev-parse HEAD)\c" >internal/version/commit

# set GOPROXY environment variable
# export GOPROXY=https://goproxy.cn

function build_release {
  # build
  go build -v -o . ./...

  # release path, for example, fsctl_go1.20.2_amd64_linux_v0.1.0
  export FSCTL_RELEASE="fsctl_${FSCTL_RELEASE_GO_VERSION}_${GOARCH}_${GOOS}_${FSCTL_RELEASE_VERSION}"

  rm -rf "$FSCTL_RELEASE"
  mkdir "$FSCTL_RELEASE"
  mv fsctl "$FSCTL_RELEASE/"

  # release archive
  tar -zcvf "$FSCTL_RELEASE.tar.gz" "$FSCTL_RELEASE"

  rm -rf "$FSCTL_RELEASE"
}

############################## linux-amd64-release ##############################

# set go env
export GOOS=linux
export GOARCH=amd64

# build
go build -v -o . ./...

FSCTL_RELEASE_GO_VERSION=$(go version | awk '{print $3}')
FSCTL_RELEASE_VERSION=$(./fsctl -v | awk 'NR==1 {print $3}')

build_release

############################## linux-amd64-release ##############################

############################## linux-arm64-release ##############################

export GOOS=linux
export GOARCH=arm64

build_release

############################## linux-arm64-release ##############################

############################# windows-release #############################

# set go env
export GOOS=windows
export GOARCH=amd64

# build
go build -v -o . ./...

# release path, for example, fsctl_go1.20.2_amd64_windows_v0.1.0
FSCTL_RELEASE="fsctl_${FSCTL_RELEASE_GO_VERSION}_${GOARCH}_${GOOS}_${FSCTL_RELEASE_VERSION}"

mkdir "$FSCTL_RELEASE"
mv fsctl.exe "$FSCTL_RELEASE/"

# release archive
zip -r "$FSCTL_RELEASE.zip" "$FSCTL_RELEASE"

rm -rf "$FSCTL_RELEASE"

############################# windows-release #############################

############################## macOS-amd64-release ##############################

export GOOS=darwin
export GOARCH=amd64

build_release

############################## macOS-amd64-release ##############################

############################## macOS-arm64-release ##############################

export GOOS=darwin
export GOARCH=arm64

build_release

############################## macOS-arm64-release ##############################

# reset commit file
echo -e "\c" >internal/version/commit

ls -alh | grep fsctl_
