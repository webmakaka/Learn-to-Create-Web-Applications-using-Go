FROM golang:1.11.5

RUN mkdir -p /project
WORKDIR /project

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/schema
RUN go get -u github.com/lib/pq
RUN go get -u github.com/jinzhu/gorm
RUN go get -u golang.org/x/crypto/bcrypt

COPY . .

CMD ["go", "run", "main.go"]