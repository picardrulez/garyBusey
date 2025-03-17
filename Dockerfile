FROM golang:1.24-bullseye
WORKDIR /app
RUN go version
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o garyBusey ./...
EXPOSE 8080
CMD ["/garyBusey"]
