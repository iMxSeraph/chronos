FROM golang:alpine AS golang
ENV LANG C.UTF-8

COPY . /go/src/muxin.io/chronos

RUN apk update && apk add git \
  && cd /go/src/muxin.io/chronos \
  && go get -u github.com/golang/dep/cmd/dep && dep ensure \
  && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go

FROM scratch
COPY --from=golang /go/src/muxin.io/chronos/main /
CMD ["/main"]
