package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

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
	userIDParam := chi.URLParam(request, "user_id")
	userID, err := strconv.Atoi(userIDParam)

	if err != nil {
		handler.log.Printf("err = %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	data, err := os.ReadFile("./database.json")
	if err != nil {
		handler.log.Printf("err = %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	var allUsersInDB Users

	if err := json.Unmarshal(data, &allUsersInDB); err != nil {
		handler.log.Printf("err = %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	var foundUser *User

	for _, user := range allUsersInDB {
		if user.ID == userID {
			foundUser = user

			break
		}
	}

	if foundUser == nil {
		handler.log.Printf("user not found = %d", userID)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	foundFriends := Users{}

	for _, userID := range foundUser.Friends {
		for _, user := range allUsersInDB {
			if user.ID == userID {
				foundFriends = append(foundFriends, user)

				break
			}
		}
	}

	response, err := json.Marshal(foundFriends)
	if err != nil {
		handler.log.Printf("user not found = %d", userID)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(response); err != nil {
		handler.log.Printf("err = %v", err)
	}
}
