package main

import (
	"curse/task59/logger"
	"net/http"

	"github.com/go-chi/chi"
)

const addr = ":8080"

func main() {
	router := chi.NewRouter()
	log := logger.New()

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
