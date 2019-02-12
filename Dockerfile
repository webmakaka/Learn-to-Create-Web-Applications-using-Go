FROM golang:1.11.5
COPY . .
RUN go run main.go