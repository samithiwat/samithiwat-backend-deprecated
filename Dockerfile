# Base Image
FROM golang:1.17.6-alpine3.15 as base

# Working directory
WORKDIR /go/app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build ./src/server.go

# Set ENV to production

ENV GO_ENV production

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./server"]