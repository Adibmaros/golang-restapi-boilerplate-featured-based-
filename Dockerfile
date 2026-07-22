# =========================================
# Stage 1: Build Stage
# =========================================
FROM golang:alpine AS builder

# Set working directory
WORKDIR /app

# Install git and ca-certificates (needed for downloading dependencies)
RUN apk add --no-cache ca-certificates git

# Copy dependency definition files first (for Docker layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build statically linked binary with optimizations (-trimpath, -s -w)
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /app/api ./cmd/api/main.go

# =========================================
# Stage 2: Minimal Final Runtime Stage
# =========================================
FROM alpine:3.21

# Add CA Certificates and Timezone data
RUN apk add --no-cache ca-certificates tzdata

# Security: Create non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary compiled from builder stage with correct ownership
COPY --from=builder --chown=appuser:appgroup /app/api /app/api

# Switch to non-root user
USER appuser

# Expose server port
EXPOSE 8080

# Healthcheck
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/swagger/index.html || exit 1

# Run application binary
CMD ["/app/api"]
