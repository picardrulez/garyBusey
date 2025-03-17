FROM golang:1.21-bullseye
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o garyBusey ./...
COPY garyBusey /garyBusey
COPY resources /resources
EXPOSE 8080
CMD ["/garyBusey"]
