# Use the same OS in both build and runtime stages
FROM golang:1.23.0 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

RUN go build -o main main.go  

# Use the same OS as the build stage (Debian-based)
FROM debian:latest

# Set working directory
WORKDIR /app/

# Copy the built binary
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Copy config file
COPY --from=builder /app/config.yaml .

# Expose the port
EXPOSE 8000

# Run the application
ENTRYPOINT ["./main"]
