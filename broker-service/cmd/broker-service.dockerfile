FROM golang:1.24.0 AS builder

WORKDIR /app

# Copy module files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest
COPY . .

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -o brokerApp ./api

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/brokerApp .
CMD ["/app/brokerApp"]