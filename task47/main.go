package main

import "net/http"

type HandlerNyan struct{}

// curl -X GET http://localhost:8080/cat-nyan
// po addressu localhost:8080/cat-nyan imeem:
//
//	writer - eto stakan kuda mii govorim
//	request - eto stakan oktuda mii slushaem
func (handler *HandlerNyan) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte("hahaha durik!")); err != nil {
		panic(err)
	}
}

type HandlerBob struct{}

// curl -X GET http://localhost:8080/cat-bob
func (handler *HandlerBob) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte("bob is tired!")); err != nil {
		panic(err)
	}
}

func main() {
	router := http.NewServeMux()

	handlerNyan := &HandlerNyan{}

	router.Handle("/cat-nyan", handlerNyan)

	handlerBob := &HandlerBob{}

	router.Handle("/cat-bob", handlerBob)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
