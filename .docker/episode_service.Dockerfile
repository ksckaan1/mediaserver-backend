FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOCACHE=/go-cache
ENV GOMODCACHE=/gomod-cache

# DOWNLOAD SERVICE DEPS
COPY ./services/episode/go.* /app/episode/
WORKDIR /app/episode
RUN --mount=type=cache,target=/gomod-cache \
    go mod download
WORKDIR /app

# DOWNLOAD SHARED DEPS
COPY ./shared/go.* /app/shared/
WORKDIR /app/shared
RUN --mount=type=cache,target=/gomod-cache \
    go mod download
WORKDIR /app

COPY . .

RUN --mount=type=cache,target=/gomod-cache \
    --mount=type=cache,target=/go-cache \
    go build -v -ldflags="-s -w" -trimpath -o /app/dist/episode /app/services/episode/cmd/api

FROM alpine:3.19

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/dist/episode /app/episode

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/episode"]