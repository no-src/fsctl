#!/usr/bin/env bash

# usage
# build the latest image
# ./scripts/build-docker.sh
# build the image with a specified tag
# ./scripts/build-docker.sh v0.0.1

# update git repository
# git pull --no-rebase

# update the latest golang image
# docker pull golang

# set the fsctl docker image name by FSCTL_IMAGE_NAME variable
FSCTL_IMAGE_NAME=nosrc/fsctl

# set the fsctl docker image tag by FSCTL_IMAGE_TAG variable
FSCTL_IMAGE_TAG=latest

# reset FSCTL_IMAGE_TAG to the value of the first parameter provided by the user
if [ -n "$1" ]; then
  FSCTL_IMAGE_TAG=$1
fi

# remove the existing old image
docker rmi -f $FSCTL_IMAGE_NAME:$FSCTL_IMAGE_TAG

# build Dockerfile
docker build --build-arg GOPROXY=$GOPROXY -t $FSCTL_IMAGE_NAME:$FSCTL_IMAGE_TAG .

# remove dangling images
docker rmi -f $(docker images -q -f "dangling=true")

# run a container to print the fsctl version
docker run -it --rm --name running-fsctl-version $FSCTL_IMAGE_NAME:$FSCTL_IMAGE_TAG fsctl -v

# push the image to the DockerHub
# docker push nosrc/fsctl:$FSCTL_IMAGE_TAG
