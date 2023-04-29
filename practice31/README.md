# Practice 31

### How to use?

```sh
GOOS=linux GOARCH=amd64 go build -o ./build/app/myapp ./server/main.go
docker-compose up -d
```

then wait for the docker infra to launch.

### Tests

In order to create an entry, do this:

```
curl -i -X POST \
  --data-binary '{"email": "second order","name": "+375 whatever"}' \
  http://localhost:8080/create
```


Do not forget to use MongoDB Compas to see the results.
