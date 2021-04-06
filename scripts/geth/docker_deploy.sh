#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

set -ex

GIT_SHORT_COMMIT=`git rev-parse --short HEAD`
TIMESTAMP=`date -u +%Y%m%d%H%M%S`
IMAGE_NAME="chainsafe/chainbridge-geth"
TAG="${TIMESTAMP}-${GIT_SHORT_COMMIT}"

docker build $BUILD_ARGS -t ${IMAGE_NAME}:${TAG} .
docker tag "${IMAGE_NAME}:${TAG}" "${IMAGE_NAME}:latest"
#echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push ${IMAGE_NAME}:latest
docker push ${IMAGE_NAME}:${TAG}