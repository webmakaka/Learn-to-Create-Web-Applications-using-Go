# Learn to Create Web Applications using Go [2018, ENG]

**Author:**  
calhoun.io

---


<br/>

## Run current version inside docker container

<br/>

```
$ docker run -it \
    -p 80:3000 \
    techhead/learn-to-create-web-applications-using-go
```

http://localhost/

<br/>

## Development

<!--
```
$ go get github.com/pilu/fresh
$ fresh
```
-->

<br/>

### lesson56

http://pgweb-demo.herokuapp.com/

<br/>

    CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        name TEXT,
        email TEXT NOT NULL
    );

    CREATE TABLE orders (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        amount INT,
        description TEXT
    )

<br/>

### lesson70

    $ cd models
    $ go test .

<br/>

### lesson 120

http://localhost:3000/galleries/new

<br/>

### lesson 140: CSRF

http://github.com/gorilla/csrf

    $ go get github.com/gorilla/csrf

<br/>

### lesson 144: Function options for services

    $ go run main.go config.go


<br/>

---

<br/>

**Marley**

Any questions in english: <a href="https://jsdev.org/chat/">Telegram Chat</a>  
Любые вопросы на русском: <a href="https://jsdev.ru/chat/">Телеграм чат</a>
