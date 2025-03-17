FROM golang:1.21-bullseye AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o garyBusey ./...
FROM debian:bullseye-slim
COPY --from-builder /app/garyBusey /garyBusey
COPY --from-builder /app/resources /resources
EXPOSE 8080
CMD["/garyBusey"]
