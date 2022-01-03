# Learn to Create Web Applications using Go [2018, ENG]

**Author:**  
calhoun.io

---


<br/>

## Run current version inside docker container


```
$ docker-compose up --build
```

<br/>

http://localhost:3000/

<br/>

```
// Create Tables
// pass pA55w0rd1
$ psql -U user1 -h localhost -p 5432 -d golang-web-demo
```

<br/>

### Development


<br/>

```sql
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
);
```

<br/>

```
golang-web-demo=# \dt
        List of relations
 Schema |  Name  | Type  | Owner 
--------+--------+-------+-------
 public | orders | table | user1
 public | users  | table | user1
(2 rows)
```


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
