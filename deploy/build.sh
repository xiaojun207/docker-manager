
DOCKER_BASE_REPO="xiaojun207"
APP_NAME="docker-manager"
DOCKER_BUILD_TAG="1.0.3"

docker build -t ${DOCKER_BASE_REPO}/${APP_NAME}:${DOCKER_BUILD_TAG} -f deploy/Dockerfile2 .
docker tag ${DOCKER_BASE_REPO}/${APP_NAME}:${DOCKER_BUILD_TAG} ${DOCKER_BASE_REPO}/${APP_NAME}:latest
docker push ${DOCKER_BASE_REPO}/${APP_NAME}:${DOCKER_BUILD_TAG}
docker push ${DOCKER_BASE_REPO}/${APP_NAME}:latest
