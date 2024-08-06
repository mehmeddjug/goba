# Use the official Go image
FROM golang:1.19-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download and install any needed dependencies
RUN go mod tidy
RUN go build -o main cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]