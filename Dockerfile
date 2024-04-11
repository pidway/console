ARG NODE_VERSION
FROM node:$NODE_VERSION as uilayer

WORKDIR /app

COPY ./web-app/package.json ./
COPY ./web-app/yarn.lock ./
RUN yarn install

COPY ./web-app .

RUN make build-static

USER node

FROM golang:1.19 as golayer

RUN apt-get update -y && apt-get install -y ca-certificates

ADD go.mod /go/src/github.com/pidway/console/go.mod
ADD go.sum /go/src/github.com/pidway/console/go.sum
WORKDIR /go/src/github.com/pidway/console

# Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download

ADD . /go/src/github.com/pidway/console
WORKDIR /go/src/github.com/pidway/console

ENV CGO_ENABLED=0

COPY --from=uilayer /app/build /go/src/github.com/pidway/console/web-app/build
RUN go build --tags=kqueue,operator -ldflags "-w -s" -a -o console ./cmd/console

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.7
MAINTAINER MinIO Development "dev@min.io"
EXPOSE 9090


COPY --from=golayer /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golayer /go/src/github.com/pidway/console .

ENTRYPOINT ["/console"]
