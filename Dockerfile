FROM golang:1.23 AS builder

# Working directory
WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Copy source code
COPY *.go .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o github-action-demo

FROM gcr.io/distroless/static-debian12

# Working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/github-action-demo .

# Hint for listening port
EXPOSE 8080/tcp

# Run the application
ENTRYPOINT ["/app/github-action-demo"]
