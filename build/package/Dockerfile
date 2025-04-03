# Use the official Golang image as the base image
FROM golang:1.24-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o app

# Create a new stage for the dev image
FROM alpine:latest AS dev

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage to the dev stage
COPY --from=build /app/app .

# Set the entrypoint command to run the application
CMD ["./app"]
