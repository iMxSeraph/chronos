FROM golang:alpine AS golang

ENV LANG C.UTF-8

WORKDIR $GOPATH/src/muxin.io/chronos
COPY . .
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./chronos"]
