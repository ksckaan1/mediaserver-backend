FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY ./services/movie/go.* ./movie/

WORKDIR /app/movie

RUN go mod download

WORKDIR /app

COPY . .

RUN go build -o ./dist/movie ./services/movie/cmd/api

FROM alpine:3.19

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/dist/movie /app/movie

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/movie"]