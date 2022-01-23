# Start golang base image
FROM golang:alpine as builder

# Install Git
RUN apk update && apk add --no-cache git

# Set the current directory inside the container
WORKDIR /app

# Copy go mod and go sum file
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code form current directory to working directory inside the container
COPY . .

# Build go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start new stage from start
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy Prebuilt binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port 8080
EXPOSE 8080

# Command to run executable
CMD ["./main"]
