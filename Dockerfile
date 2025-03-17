FROM golang:1.24-bullseye
WORKDIR /app
RUN go version
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN ls -l /app
#RUN go build -o garyBusey ./... || echo "go build failed" && exit 1
#COPY /app/garyBusey /garyBusey
#COPY resources /resources
#EXPOSE 8080
#CMD ["/garyBusey"]
