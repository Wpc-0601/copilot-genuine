# syntax=ubuntu-k8s-1.23.sh/dockerfile:1
# Use the official golang image as a builder stage
# This image contains the Go compiler and tools
FROM golang:1.20 as builder

# Set the working directory inside the container
# This is where the source code and the executable will be stored
WORKDIR /app

# Copy the source code from the current directory to the working directory
# This assumes that your main.go file is in the same directory as the Dockerfile
COPY src/ .

# Build the application with CGO disabled and static linking
# CGO is a feature that allows Go to call C code, but it can cause compatibility issues
# Static linking means that all dependencies are included in the executable
# The -s and -w flags strip debugging information to reduce the executable size
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main .

# Use alpine as a base image for the final stage
# Alpine is a minimal Linux distribution that is only 5 MB in size
FROM alpine

# Set the working directory inside the container
# This is where the executable will be copied to
WORKDIR /root/

# Copy the executable from the builder stage
# This uses a special syntax to copy files from another image
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
# This tells Docker that the container listens on port 8080
EXPOSE 8080

# Run the executable
# This is the command that will be executed when the container starts
CMD ["./main"]