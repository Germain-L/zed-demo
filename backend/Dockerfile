# Use the official Golang image as the base image
FROM golang:1.24.3-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application's source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Start a new stage from scratch
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the 'builder' stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
