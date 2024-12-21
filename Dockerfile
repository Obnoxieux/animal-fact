# Stage 1: Build the Go application
FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a minimal container for the app
FROM ubuntu:latest
RUN apt install mc
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/assets ./assets

EXPOSE 8090

CMD ["./main"]
