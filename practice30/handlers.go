package main

import (
	"errors"
	"net/http"
)

type User struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type Users []*User

type MakeNewFriendInfo struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}

var (
	ErrNotFound = errors.New("not found")
)

type HandlerUpdateUserAge struct {
	log Logger
}

type HandlerDeleteUser struct {
	log Logger
}

func (handler *HandlerDeleteUser) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"target_id":"1"}`)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}
