FROM golang:1.12 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

FROM golang:1.12
COPY --from=build-env /go/bin/app /
CMD ["/app"]
