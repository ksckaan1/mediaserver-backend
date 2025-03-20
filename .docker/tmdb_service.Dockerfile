FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY ./tmdb_service/go.* .

WORKDIR /app/tmdb_service

RUN go mod download

WORKDIR /app

COPY . .

RUN go build -o ./dist/tmdb_service ./tmdb_service/cmd/api

FROM alpine:3.19

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/dist/tmdb_service /app/tmdb_service

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/tmdb_service"]