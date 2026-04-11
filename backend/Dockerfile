# Stage 1: Build stage
FROM golang:1.26.2-alpine3.23 AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY backend/*.go ./
COPY frontend ./frontend

# Build the application
# CGO_ENABLED=0 for a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final stage (Production)
FROM alpine:3.23.3

# Install CA certificates for secure connections
RUN apk --no-cache add ca-certificates

# Create a non-root user first
RUN adduser -D -u 1000 appuser

WORKDIR /app

# Copy the binary and frontend from builder with correct ownership
COPY --from=builder --chown=appuser:appuser /app/main .
COPY --from=builder --chown=appuser:appuser /app/frontend ./frontend

USER appuser

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
