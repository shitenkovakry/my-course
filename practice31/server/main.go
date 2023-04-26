package main

import (
	"curse/practice31/datasource/mongo"
	"curse/practice31/datasource/users"
	handler_create "curse/practice31/handlers/user/create"
	"curse/practice31/logger"
	"net/http"

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
