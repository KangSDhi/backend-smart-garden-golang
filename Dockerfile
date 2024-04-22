FROM golang:1.22.2 AS builder
LABEL authors="kangsdhi"

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start from scratch (no base image)
FROM scratch

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV ZONEINFO=/zoneinfo.zip
COPY .env.prod /app/.env

EXPOSE 3000

CMD ["/app/main"]