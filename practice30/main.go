package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type Logger interface {
	Printf(format string, v ...any)
}

func main() {
	router := chi.NewRouter()

	logger := log.New(os.Stdout, "practice30", log.Flags())

	handlerCreateUser := &HandlerCreateUser{logger}
	router.Method(http.MethodPost, "/create", handlerCreateUser)

	handlerMakeFriends := &HandlerMakeFriends{logger}
	router.Method(http.MethodPost, "/make_friends", handlerMakeFriends)

	handlerDeleteUser := &HandlerDeleteUser{logger}
	router.Method(http.MethodDelete, "/user", handlerDeleteUser)

	handlerReturnAllFriendsID := &HandlerReturnAllFriendsID{logger}
	router.Method(http.MethodGet, "/friends/{user_id}", handlerReturnAllFriendsID)

	handlerUpdateUserAge := &HandlerUpdateUserAge{logger}
	router.Method(http.MethodPut, "/{user_id}", handlerUpdateUserAge)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
