FROM golang:1.23.5 AS builder

WORKDIR /app

# Copy the go.work file (if using Go workspace) and download dependencies
COPY . .

WORKDIR /app/cmd
RUN go work sync && CGO_ENABLED=0 GOOS=linux go build -o /app/cmd main.go

FROM alpine:latest

RUN apk add --no-cache nginx && rm -rf /var/cache/apk/*

# Set up a working directory in the runtime container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/ .

# Expose the port your service will run on (adjust if needed)
EXPOSE 5000

# Command to run the nginx service binary
CMD ["sh", "-c", "/app/cmd/main"]