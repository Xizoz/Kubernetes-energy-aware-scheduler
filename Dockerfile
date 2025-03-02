# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary with CGO disabled (static binary)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o custom-scheduler main.go

# Use a minimal base image for execution
FROM scratch 

# Set working directory
WORKDIR /root/

# Copy the compiled binary from builder stage
COPY --from=builder /app/custom-scheduler .

# Run the scheduler
CMD ["/root/custom-scheduler"]
