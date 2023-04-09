package main

import (
	"log"
	"net/http"
	"os"

	"curse/task58/handlers"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	logger := log.New(os.Stdout, "task58", log.Flags())

	handlerForPlayAsX := handlers.NewHandlerForPlayingX(logger)
	router.Method(http.MethodPost, "api/v1/mark/x", handlerForPlayAsX)

	handlerForPlayAsO := handlers.NewHandlerForPlayingY(logger)
	router.Method(http.MethodPost, "/api/v1/mark/o", handlerForPlayAsO)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Print("serving at :8080")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
