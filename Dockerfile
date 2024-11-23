# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy all files into the container's /app directory
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -o myapp

# Stage 2: Create the minimal runtime environment
FROM alpine:latest

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/myapp /

# Set an environment variable to configure the port (default to 8080)
ENV PORT=8080

# Expose the port the application will listen on
EXPOSE 8080

# Command to run the Go application (listening on the port specified)
ENTRYPOINT ["/myapp"]
