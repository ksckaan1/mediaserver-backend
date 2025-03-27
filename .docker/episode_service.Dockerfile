FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY ./services/episode/go.* ./episode/

WORKDIR /app/episode

RUN go mod download

WORKDIR /app

COPY . .

RUN go build -o ./dist/episode ./services/episode/cmd/api

FROM alpine:3.19

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/dist/episode /app/episode

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/episode"]