FROM golang:1.21-bullseye
WORKDIR /app
RUN go version
COPY go.mod go.sum ./
RUN ls -l /app
#RUN go mod download || echo "Go module download failed" && exit 1
#COPY . .
#RUN go build -o garyBusey ./... || echo "go build failed" && exit 1
#COPY /app/garyBusey /garyBusey
#COPY resources /resources
#EXPOSE 8080
#CMD ["/garyBusey"]
