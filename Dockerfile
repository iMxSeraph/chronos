FROM golang:alpine AS golang

ENV LANG C.UTF-8

WORKDIR $GOPATH/src/muxin.io/chronos
COPY . .
RUN apk update && apk add git && go get -u github.com/golang/dep/cmd/dep && dep ensure && go build .

EXPOSE 8080
ENTRYPOINT ["./chronos"]
