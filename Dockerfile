FROM golang:1.22.5-alpine

# Set working directory and copy all files to it.
WORKDIR /app
COPY . .

# Download all dependencies and build a binary.
RUN go mod download
RUN go build -o main ./cmd/gin-pgx-api

# Run binary with specified interface and multicast address.
CMD ["./main", "./cmd/gin-pgx-api/config.json"]
