VERSION=$(git describe --tags `git rev-list --tags --max-count=1`)
VERSION=${VERSION/v/} # eg.: 0.2.5
BUILD_DATE=`date +%Y%m%d` # eg.: 20220701
COMMIT_HASH=`git rev-parse HEAD` #

echo "$VERSION"
# CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o App App.go
go build -ldflags "-s -w -X docker-manager/service.Version=$VERSION -X docker-manager/service.Date=$BUILD_DATE -X docker-manager/service.CommitHash=$COMMIT_HASH" -o App App.go
