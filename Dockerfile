FROM golang:1.11.5

RUN mkdir -p /project
WORKDIR /project

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/schema
RUN go get github.com/lib/pq
RUN go get github.com/jinzhu/gorm

COPY . .

CMD ["go", "run", "main.go"]