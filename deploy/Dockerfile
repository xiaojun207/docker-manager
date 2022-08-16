#build源镜像
FROM golang:1.17.4 as build
#作者
MAINTAINER xiaojun "xiaojun207@126.com"

#ENV GOPROXY https://goproxy.io
ENV GO111MODULE on
ARG VERSION
ARG COMMIT_HASH
ARG BUILD_DATE

WORKDIR /go/release

ADD go.mod .
ADD go.sum .
RUN go mod download
RUN go mod tidy

WORKDIR /go/release

ADD . .

#RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o App App.go
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w  -X docker-manager/service.Version=$VERSION -X docker-manager/service.Date=$BUILD_DATE -X docker-manager/service.CommitHash=$COMMIT_HASH" -o App App.go

#运行镜像
FROM alpine:latest AS production

RUN mkdir /app
WORKDIR /app

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/App /app/
COPY --from=build /go/release/views /app/views
COPY --from=build /go/release/deploy/entrypoint.sh /app/
COPY --from=build /go/release/config.yml /app/
COPY --from=build /go/release/config-example-mysql.yml /app/
COPY --from=build /go/release/config-example-sqlite3.yml /app/

ENV driverName=""\
    useCache=""\
    consoleCode=""\
    dataSourceUrl=""

ENTRYPOINT ["sh", "/app/entrypoint.sh"]

#暴露端口
EXPOSE 8068




