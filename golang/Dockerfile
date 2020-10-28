FROM golang:1.12 as build-env
ENV GO111MODULE=on

WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
ADD . .
RUN go build -o /go/bin/app

FROM ubuntu:bionic
COPY --from=build-env /go/bin/app /
CMD ["/app"]
