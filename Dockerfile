# Stage 1: Build
FROM golang:1.23.4-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o message-broker cmd/main.go .

# Stage 2: Run
FROM alpine:3.21

# Install necessary CA certificates
RUN apk --no-cache add ca-certificates

# Set working directory and copy the compiled binary
WORKDIR /root/
COPY --from=builder /app/message-broker .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./message-broker"]