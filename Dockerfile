FROM golang:1.15.5

WORKDIR /app

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/schema
RUN go get -u github.com/lib/pq
RUN go get -u github.com/jinzhu/gorm
RUN go get -u golang.org/x/crypto/bcrypt
RUN go get -u github.com/gorilla/csrf

COPY ./app ./

CMD ["go", "run", "main.go", "config.go"]
