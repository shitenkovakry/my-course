package main

import (
	"log"
	"net/http"
	"os"
)

type Logger interface {
	Printf(format string, v ...any)
}

func main() {
	router := http.NewServeMux()

	logger := log.New(os.Stdout, "practice30", log.Flags())

	handlerUpdateUserAge := &HandlerUpdateUserAge{logger}
	router.Handle("/user_id", handlerUpdateUserAge)

	handlerCreateUser := &HandlerCreateUser{logger}
	router.Handle("/create", handlerCreateUser)

	handlerMakeFriends := &HandlerMakeFriends{logger}
	router.Handle("/make_friends", handlerMakeFriends)

	handlerDeleteUser := &HandlerDeleteUser{logger}
	router.Handle("/user", handlerDeleteUser)

	handlerReturnAllFriendsID := &HandlerReturnAllFriendsID{}
	router.Handle("/friends/user_id", handlerReturnAllFriendsID)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
