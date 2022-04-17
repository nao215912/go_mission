FROM golang:1.17-alpine3.14 as build


WORKDIR /go/api

COPY . .

ARG AIR_VERSION=v1.27.3
ARG DLV_VERSION=v1.7.1

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl make && \
  go install github.com/cosmtrek/air@${AIR_VERSION} && \
  go install github.com/go-delve/delve/cmd/dlv@${DLV_VERSION} && \
  go install golang.org/x/tools/cmd/goimports@latest

RUN set -eux && \
  go build -o go_mission ./api/main.go

ENV CGO_ENABLED 0

FROM alpine:3.14

WORKDIR /api

COPY --from=build /go/api/go_mission .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /api/go_mission

CMD ["./go_mission"]