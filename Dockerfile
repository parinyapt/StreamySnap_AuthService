# syntax=docker/dockerfile:1

FROM golang:1.20 AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux

WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build
RUN go build -o ./docker-streamysnap-authservice ./cmd

EXPOSE 80

# Run
CMD ["./docker-streamysnap-authservice", "-mode=production"]