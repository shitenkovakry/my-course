package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"curse/practice31/datasource/mongo"
	"curse/practice31/datasource/users"
	handler_create "curse/practice31/handlers/user/create"
	handler_delete "curse/practice31/handlers/user/delete"
	handler_make_friend "curse/practice31/handlers/user/make-friend"
	handler_update_user "curse/practice31/handlers/user/update-user"
	"curse/practice31/logger"

	"github.com/go-chi/chi"
)

const (
	addr = ":8080"
)

func main() {
	router := chi.NewRouter()
	log := logger.New()
	mongoDB := mongo.New(log, "", "", []string{"mongodb:27017"}, "my-database")
	usersManager := users.New(log, mongoDB)

	handlerForCreateUser := handler_create.NewHandlerForCreateUser(log, usersManager)
	router.Method(http.MethodPost, "/create", handlerForCreateUser)

	handlerForUpdateUserByAge := handler_update_user.NewHandlerForUpdateUserByAge(log, usersManager)
	router.Method(http.MethodPut, "/{user_id}", handlerForUpdateUserByAge)

	handlerForDeleteUser := handler_delete.NewHandlerForDeleteUser(log, usersManager)
	router.Method(http.MethodDelete, "/user", handlerForDeleteUser)

	handlerForMakeFriend := handler_make_friend.NewHandlerForMakeFriend(log, usersManager)
	router.Method(http.MethodPost, "/make_friends", handlerForMakeFriend)

	server := NewServer(addr, router)

	log.Printf("Serving at [%s]", addr)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server is error: %v", err)
		}
	}()

	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// завершение работы серверов
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server1 shutdown error: %v", err)
	}
}

func NewServer(address string, router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
