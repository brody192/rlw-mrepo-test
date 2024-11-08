# Stage 1: Builder
FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY ./pkg ./pkg
COPY ./service ./service

WORKDIR /app/service

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go

# Stage 2: Final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/main ./

CMD ["./main"]
