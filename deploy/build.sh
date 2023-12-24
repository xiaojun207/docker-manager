#!/bin/bash
# author xiaojun207

DOCKER_BASE_REPO="xiaojun207"
APP_NAME="docker-manager"

VERSION=$(git describe --tags `git rev-list --tags --max-count=1`)
VERSION=${VERSION/v/} # eg.: 1.5.6
DOCKER_BUILD_TAG="${VERSION}"
BUILD_DATE=`date +%Y%m%d` # eg.: 20220701
COMMIT_HASH=`git rev-parse HEAD` #

docker build \
  --build-arg VERSION="${VERSION}" \
  --build-arg COMMIT_HASH="${COMMIT_HASH}" \
  --build-arg BUILD_DATE="${BUILD_DATE}" \
   -t ${DOCKER_BASE_REPO}/${APP_NAME}:${VERSION} \
   -f deploy/Dockerfile .

#docker tag ${DOCKER_BASE_REPO}/${APP_NAME}:${VERSION} ${DOCKER_BASE_REPO}/${APP_NAME}:latest
#docker push ${DOCKER_BASE_REPO}/${APP_NAME}:${VERSION}
#docker push ${DOCKER_BASE_REPO}/${APP_NAME}:latest
