# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the project files into the container
COPY . .

# Build the project inside the container
RUN go build -o main .

# Expose the port that the application listens on
EXPOSE 8080

# Set the command to run the application
CMD ["./main"]