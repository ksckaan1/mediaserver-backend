FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOCACHE=/go-cache
ENV GOMODCACHE=/gomod-cache

# DOWNLOAD SERVICE DEPS
COPY ./services/apidoc/go.* /app/apidoc/
WORKDIR /app/apidoc
RUN --mount=type=cache,target=/gomod-cache \
    go mod download
WORKDIR /app

COPY . .

RUN --mount=type=cache,target=/gomod-cache \
    --mount=type=cache,target=/go-cache \
    go build -v -ldflags="-s -w" -trimpath -o /app/dist/apidoc /app/services/apidoc

FROM alpine:3.19

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/dist/apidoc /app/apidoc

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/apidoc"]