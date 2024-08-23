# Dockerfile for users service
FROM golang:1.20.0-alpine

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum files to the working directory.
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the application.
RUN go build -o main .

# Expose the application port (e.g., 8080)
EXPOSE 8080

# Run the executable.
CMD ["./main"]