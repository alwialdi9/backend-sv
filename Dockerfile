FROM golang:1.22.3 AS builder

# Set the working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .
COPY ./cmd/.env ./

# Build the application, specifying cmd/main.go
RUN GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go

# Deployment Stage
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .
COPY cmd/.env .env

# Ensure the binary is executable
RUN chmod +x /root/main

EXPOSE 3000

# Run the binary
ENTRYPOINT ["/root/main"]