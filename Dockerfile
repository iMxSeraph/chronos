FROM golang:alpine AS golang

ENV LANG C.UTF-8

WORKDIR $GOPATH/src/muxin.io/chronos
COPY . .
RUN apk update && apk add curl && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && dep ensure && go build .

EXPOSE 8080
ENTRYPOINT ["./chronos"]
