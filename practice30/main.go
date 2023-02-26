package main

import "net/http"

type HandlerUpdateUserAge struct{}

func (handler *HandlerUpdateUserAge) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"new age":"28"}`)); err != nil {
		panic(err)
	}
}

type HandlerCreateUser struct{}

func (handler *HandlerCreateUser) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"name":"some name","age":"24","friends":[]}`)); err != nil {
		panic(err)
	}
}

type HandlerMakeFriends struct{}

func (handler *HandlerMakeFriends) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"source_id":"1","target_id":"2"}`)); err != nil {
		panic(err)
	}
}

type HandlerDeleteUser struct{}

func (handler *HandlerDeleteUser) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"target_id":"1"}`)); err != nil {
		panic(err)
	}
}

type HandlerReturnAllFriendsID struct{}

func (handler *HandlerReturnAllFriendsID) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(``)); err != nil {
		panic(err)
	}
}

func main() {
	router := http.NewServeMux()

	handlerUpdateUserAge := &HandlerUpdateUserAge{}
	router.Handle("/user_id", handlerUpdateUserAge)

	handlerCreateUser := &HandlerCreateUser{}
	router.Handle("/create", handlerCreateUser)

	handlerMakeFriends := &HandlerMakeFriends{}
	router.Handle("/make_friends", handlerMakeFriends)

	handlerDeleteUser := &HandlerDeleteUser{}
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
