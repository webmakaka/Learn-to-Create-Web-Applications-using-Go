# Learn to Create Web Applications using Go [2018, ENG]

calhoun.io

---

    $ docker run -i -t -p 3000:3000 marley/learn-to-create-web-applications-using-go

    \$ go get github.com/pilu/fresh
    \$ fresh

    \$ go get github.com/gorilla/mux

<br/>

### lesson56

http://pgweb-demo.herokuapp.com/

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

### Run inside docker container

```
$ docker run -it \
    -p 80:3000 \
    techhead/react-practice-course-project2
```


---

**Marley**

<a href="https://marley.org">marley.org</a>
