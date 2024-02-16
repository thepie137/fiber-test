# Use the official Golang image with Alpine Linux
FROM golang:1.22.0-alpine3.18 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:3.18

# Create a directory to store the application binary
RUN mkdir /app

# Copy the pre-built binary from the previous stage to the /app directory in the new stage
COPY --from=build /app/main /app/main

# Expose port 8005 to the outside world
EXPOSE 8005

# Command to run the executable
CMD ["/app/main"]
