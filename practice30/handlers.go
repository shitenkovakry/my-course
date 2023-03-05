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

type DeleteUserInfo struct {
	ID int `json:"target_id"`
}

var (
	ErrNotFound = errors.New("not found")
)

func FindUserByID(allUsersInDB Users, id int) (*User, error) {
	for i := 0; i <= len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID == id {
			return user, nil
		}
	}

	return nil, ErrNotFound
}

type HandlerUpdateUserAge struct {
	log Logger
}

func (handler *HandlerUpdateUserAge) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"new age":"28"}`)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}
