#build源镜像
FROM golang:1.16.5 as build
#作者
MAINTAINER xiaojun "xiaojun207@126.com"

#ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o App App.go

#运行镜像
FROM alpine:latest AS production

RUN mkdir /app
WORKDIR /app

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/App /app/
COPY --from=build /go/release/views /app/views
COPY --from=build /go/release/deploy/entrypoint.sh /app/

ENV driverName="mysql"\
    dataSourceUrl="root:Abc123@(dbhost:3306)/docker-manager?charset=utf8"

ENTRYPOINT ["sh", "/app/entrypoint.sh"]

#暴露端口
#EXPOSE 8068




