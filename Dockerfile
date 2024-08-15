# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o /translate-app

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the Go app
CMD ["/translate-app"]
