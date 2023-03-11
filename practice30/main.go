package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"curse/practice30/handlers"
)

func main() {
	router := chi.NewRouter()

	logger := log.New(os.Stdout, "practice30", log.Flags())

	logger.Print("durik nachinaet")

	handlerCreateUser := handlers.NewHandlerCreateUser(logger)
	router.Method(http.MethodPost, "/create", handlerCreateUser)

	handlerMakeFriends := handlers.NewHandlerMakeFriends(logger)
	router.Method(http.MethodPost, "/make_friends", handlerMakeFriends)

	handlerDeleteUser := handlers.NewHandlerDeleteUser(logger)
	router.Method(http.MethodDelete, "/user", handlerDeleteUser)

	handlerReturnAllFriendsID := handlers.NewHandlerReturnAllFriendsID(logger)
	router.Method(http.MethodGet, "/friends/{user_id}", handlerReturnAllFriendsID)

	handlerUpdateUserAge := handlers.NewHandlerUpdateUserAge(logger)
	router.Method(http.MethodPut, "/{user_id}", handlerUpdateUserAge)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Print("serving at :8080")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
