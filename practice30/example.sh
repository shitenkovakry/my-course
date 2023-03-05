curl -i -X POST --data-raw '{"email":"kuzia@mail.ru","name":"some name","age":24,"friends":[]}' http://localhost:8080/create

curl -i -X POST http://localhost:8080/make_friends --data-raw '{"source_id":34,"target_id":33}'

curl -i -X DELETE http://localhost:8080/user --data-raw '{"target_id":12}'

curl -i -X DELETE http://localhost:8080/user --data-raw '{"target_id":3}'

curl -i -X PUT http://localhost:8080/13 --data-raw '{"new_age":28}'
