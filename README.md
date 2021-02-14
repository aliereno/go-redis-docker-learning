# Go Server with Redis, Docker

* go 1.15+
* [gofiber/fiber](https://github.com/gofiber/fiber)
* [go-redis/redis](https://github.com/go-redis/redis)

### Installation & Run:

On terminal;

```
cp sample.env .env
docker-compose up
```

## Country API

#### /:alpha2letter

* `GET` Body:{ "name": STRING, "code": STRING }

###EXAMPLE
* `http://127.0.0.1:3000/TR` -> `{"name":"Egypt","code":"EG"}`
* `http://127.0.0.1:3000/GB` -> `{"name":"United Kingdom","code":"GB"}`
