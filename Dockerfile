# Start with a minimal Go image
FROM golang:1.21 as builder

# Set working directory
WORKDIR /app

# Copy and build the custom scheduler
COPY . .
RUN go mod tidy && go build -o custom-scheduler main.go

# Use a minimal base image for execution
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary from builder
COPY --from=builder /app/custom-scheduler .

# Run the scheduler
CMD ["./custom-scheduler"]
