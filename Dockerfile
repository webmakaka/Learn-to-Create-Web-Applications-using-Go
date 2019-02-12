FROM golang:1.11.5

RUN go get github.com/gorilla/mux

COPY . .
RUN go run main.go