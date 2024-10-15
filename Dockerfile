# Use Golang official image as the base image
FROM golang:1.23-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum don't change
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the Go app
CMD ["./main"]