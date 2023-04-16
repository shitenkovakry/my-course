package main

import (
	"curse/task59/datasource/users"
	handler_create_user "curse/task59/handlers/create-user"

	handler_delete_user "curse/task59/handlers/delete-user"
	handler_make_friend "curse/task59/handlers/make-friend"
	handler_update_user "curse/task59/handlers/update-user"
	"curse/task59/logger"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	addr = ":8080"
)

func main() {
	router := chi.NewRouter()
	log := logger.New()
	users := users.New()

	handlerCreateUser := handler_create_user.NewHandlerForCreateUser(log, users)
	router.Method(http.MethodPost, "/create", handlerCreateUser)

	handlerUpdateUserByAge := handler_update_user.NewHandlerForUpdateUserByAge(log, users)
	router.Method(http.MethodPut, "/{user_id}", handlerUpdateUserByAge)

	handlerDeleteUser := handler_delete_user.NewHandlerForDeleteUser(log, users)
	router.Method(http.MethodDelete, "/user", handlerDeleteUser)

	handlerMakeFriend := handler_make_friend.NewHandlerForMakeFriend(log, users)
	router.Method(http.MethodPost, "/make_friends", handlerMakeFriend)

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
