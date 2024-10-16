# Stage 1: Build the Go app
FROM golang:1.23-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum don't change
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the Go app
RUN go build -o /bin/app ./cmd/main.go

# Stage 2: Run the Go app
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/app /bin/app

# Expose the port the app runs on
EXPOSE 8080

# Command to run the Go app
CMD ["/bin/app"]