ARG GOLANG_VERSION

# First stage: create certificates.
FROM alpine:latest as certs-builder
RUN apk --update add ca-certificates

# Second stage: setup go-build.
FROM golang:${GOLANG_VERSION}-alpine AS go-builder

WORKDIR /app

ENV GO111MODULE=on\
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Third stage: download go module dependencies.
COPY ./go.* ./
RUN go mod download
COPY ./ ./

# Fourth stage: build go binary.
RUN go build \
  -installsuffix 'static' \
  -o start ./cmd/app

# Fifth stage: setup permissions, ssl certs, and server executable.
FROM scratch AS final

COPY --from=certs-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /app/start /app/start

CMD ["/app/start"]
