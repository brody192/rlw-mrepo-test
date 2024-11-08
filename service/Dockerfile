# Stage 1: Builder
FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY ../pkg /app/pkg
RUN ls -la /app/pkg
COPY . /app/service
RUN ls -la /app/service

WORKDIR /app/service


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go

# Stage 2: Final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
ENV PORT=8080
EXPOSE 8080
CMD ["./main"]
