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


In order to update an entry, do this:

```
curl -i -X PUT \
  --data-binary '{"age": 39}' \
  http://localhost:8080/3
```


In order to delete an entry, do this:

```
curl -i -X DELETE \
  --data-binary '{"id": 3}' \
  http://localhost:8080/user
```


In order to make friends an entry, do this:

```
curl -i -X POST \
  --data-binary '{"source_id": 3, "target_id":2}' \
  http://localhost:8080/make_friends
```

Do not forget to use MongoDB Compas to see the results.
