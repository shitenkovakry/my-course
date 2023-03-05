package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

type HandlerUpdateUserAge struct {
	log Logger
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

func (handler *HandlerUpdateUserAge) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"new age":"28"}`)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}

type HandlerMakeFriends struct {
	log Logger
}

func (handler *HandlerMakeFriends) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		makeNewFriend *MakeNewFriendInfo
		allUsersInDB  Users
	)

	dataReadAllUsersInDB, err := os.ReadFile("./database.json")
	if err != nil {
		handler.log.Printf("err = %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.Unmarshal(dataReadAllUsersInDB, &allUsersInDB); err != nil {
		handler.log.Printf("err = %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.Unmarshal(body, &makeNewFriend); err != nil {
		handler.log.Printf("cannot unmarshal body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	sourceUser, err := FindUserByID(allUsersInDB, makeNewFriend.SourceID)
	if err != nil {
		handler.log.Printf("cannot find user %d", makeNewFriend.SourceID)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	targetUser, err := FindUserByID(allUsersInDB, makeNewFriend.TargetID)
	if err != nil {
		handler.log.Printf("cannot find user %d", makeNewFriend.TargetID)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	sourceUser.Friends = append(sourceUser.Friends, targetUser.ID)
	targetUser.Friends = append(targetUser.Friends, sourceUser.ID)

	dataWriteAllUsersInDB, err := json.Marshal(allUsersInDB)
	if err != nil {
		handler.log.Printf("cannot marshal allUsersInDB")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := os.WriteFile("./database.json", dataWriteAllUsersInDB, os.ModePerm); err != nil {
		handler.log.Printf("cannot write allUsersInDB in file")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	response := fmt.Sprintf("%s and %s are friends now", sourceUser.Name, targetUser.Name)

	if _, err := writer.Write([]byte(response)); err != nil {
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
