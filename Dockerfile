FROM golang:1.11.5

RUN mkdir -p /project
WORKDIR /project

RUN go get github.com/gorilla/mux

COPY . .

CMD ["go", "run", "main.go"]