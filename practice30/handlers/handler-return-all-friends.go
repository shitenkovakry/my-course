package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

type HandlerReturnAllFriendsID struct {
	log Logger
}

func NewHandlerReturnAllFriendsID(log Logger) *HandlerReturnAllFriendsID {
	return &HandlerReturnAllFriendsID{
		log: log,
	}
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
