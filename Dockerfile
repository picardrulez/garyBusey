FROM golang:1.24-bullseye
WORKDIR /app
RUN go version
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o garyBusey ./...
RUN ls -l /app
#COPY /app/garyBusey /garyBusey
#COPY resources /resources
#EXPOSE 8080
#CMD ["/garyBusey"]
