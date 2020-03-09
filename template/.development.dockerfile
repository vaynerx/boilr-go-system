ARG GOLANG_VERSION

# First stage: create certificates.
FROM alpine:latest as certs-builder
RUN apk --update add ca-certificates

# Second stage: download go module deps.
FROM golang:${GOLANG_VERSION}-alpine AS go-builder

WORKDIR /app

ENV GO111MODULE=on\
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY ./go.* ./
RUN go mod download
COPY ./ ./

# Third stage: download Compile Daemon, build go binary, and run it.
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o start ./cmd/app/app.go" -command="./start"