FROM golang:1.16-alpine as builder

RUN mkdir /build

COPY . /build/go-www-geotag

RUN apk update && apk upgrade \
    && apk add make libc-dev gcc git \
    && cd /build/go-www-geotag \
    && go build -mod vendor -o /usr/local/bin/geotag-www cmd/server/main.go    

FROM alpine:latest

COPY --from=builder /usr/local/bin/geotag-www /usr/local/bin/

RUN mkdir /usr/local/data

# See notes in data/README.md

COPY data/*.db /usr/local/data
COPY data/*.pmtiles /usr/local/data

RUN apk update && apk upgrade \
    && apk add ca-certificates

ENTRYPOINT ["/usr/local/bin/geotag-www"]