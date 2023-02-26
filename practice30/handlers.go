package main

import "net/http"

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type Users []*User

type HandlerUpdateUserAge struct {
	log Logger
}

func (handler *HandlerUpdateUserAge) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"new age":"28"}`)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}

type HandlerCreateUser struct {
	log Logger
}

func (handler *HandlerCreateUser) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"name":"some name","age":"24","friends":[]}`)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}

type HandlerMakeFriends struct {
	log Logger
}

func (handler *HandlerMakeFriends) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"source_id":"1","target_id":"2"}`)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}

type HandlerDeleteUser struct {
	log Logger
}

func (handler *HandlerDeleteUser) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"target_id":"1"}`)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}

type HandlerReturnAllFriendsID struct {
	log Logger
}

func (handler *HandlerReturnAllFriendsID) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(``)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}
