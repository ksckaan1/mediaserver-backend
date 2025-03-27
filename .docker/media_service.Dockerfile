FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY ./services/media/go.* ./media/

WORKDIR /app/media

RUN go mod download

WORKDIR /app

COPY . .

RUN go build -o ./dist/media ./services/media/cmd/api

FROM alpine:3.19

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/dist/media /app/media

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/media"]