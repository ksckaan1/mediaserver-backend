FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY ./series_service/go.* .

WORKDIR /app/series_service

RUN go mod download

WORKDIR /app

COPY . .

RUN go build -o ./dist/series_service ./series_service/cmd/api

FROM alpine:3.19

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/dist/series_service /app/series_service

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/series_service"]