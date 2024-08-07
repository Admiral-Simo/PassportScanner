# Use the official Golang image to build the application
FROM golang:1.22.4

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the entire source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Expose port 4000 to the outside world
EXPOSE 4000

# Command to run the executable
CMD ["./main"]
