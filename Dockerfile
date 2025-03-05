# Build stage with Go 1.24
FROM golang:1.24 AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o energy-scheduler main.go

# Final minimal execution image
FROM scratch
WORKDIR /root/
COPY --from=builder /app/energy-scheduler .
CMD ["/root/energy-scheduler"]
