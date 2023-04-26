package main

import (
	"net/http"

	"curse/practice31/datasource/mongo"
	"curse/practice31/datasource/users"
	handler_create "curse/practice31/handlers/user/create"
	handler_update_user "curse/practice31/handlers/user/update-user"
	"curse/practice31/logger"

	"github.com/go-chi/chi"
)

const addr = ":8080"

func main() {
	router := chi.NewRouter()
	log := logger.New()
	mongoDB := mongo.New(log, "", "", []string{"localhost:27017"}, "my-database")
	usersManager := users.New(log, mongoDB)

	handlerForCreateUser := handler_create.NewHandlerForCreateUser(log, usersManager)
	router.Method(http.MethodPost, "/create", handlerForCreateUser)

	handlerForUpdateUserByAge := handler_update_user.NewHandlerForUpdateUserByAge(log, usersManager)
	router.Method(http.MethodPut, "/{user_id}", handlerForUpdateUserByAge)

	server := NewServer(addr, router)

	log.Printf("Serving at [%s]", addr)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func NewServer(address string, router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
