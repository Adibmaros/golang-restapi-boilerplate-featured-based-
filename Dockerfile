# =========================================
# Stage 1: Build Stage
# =========================================
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Install git and ca-certificates (needed for downloading dependencies)
RUN apk add --no-cache ca-certificates git

# Copy dependency definition files first (for Docker layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build statically linked binary with optimizations (-s -w to remove debug info and strip symbol table)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/api ./cmd/api/main.go

# =========================================
# Stage 2: Minimal Final Runtime Stage
# =========================================
FROM alpine:3.21

# Add CA Certificates and Timezone data
RUN apk add --no-cache ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy binary compiled from builder stage
COPY --from=builder /app/api /app/api

# Expose server port
EXPOSE 8080

# Run application binary
CMD ["/app/api"]
