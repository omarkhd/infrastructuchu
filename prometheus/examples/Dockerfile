FROM golang:1.13.4-alpine3.10

MAINTAINER omarkhd.mx@gmail.com

WORKDIR /go/src/omarkhd

RUN apk add git
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promauto
RUN go get github.com/prometheus/client_golang/prometheus/promhttp

ADD . .
RUN go build -o examples

ENTRYPOINT ./examples

