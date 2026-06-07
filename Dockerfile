# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install dependencies
# RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g ./cmd/main.go

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Run stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary and migrations
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/docs ./docs
COPY .env .

# Expose port
EXPOSE 8000

# Run the application
CMD ["./main"]