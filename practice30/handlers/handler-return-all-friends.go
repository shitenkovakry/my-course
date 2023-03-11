package handlers

import (
	"encoding/json"
	"net/http"
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

func FindAllFrindsByUserID(allUsersInDB Users, userID int) (Users, error) {
	var foundUser *User

	for _, user := range allUsersInDB {
		if user.ID == userID {
			foundUser = user

			break
		}
	}

	if foundUser == nil {
		return nil, ErrNotFound
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

	return foundFriends, nil
}

func (handler *HandlerReturnAllFriendsID) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userIDParam := chi.URLParam(request, "user_id")
	userID, err := strconv.Atoi(userIDParam)

	if err != nil {
		handler.log.Printf("err = %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	allUsersInDB, err := ReadAllUsers()
	if err != nil {
		handler.log.Printf("cannot unmarshal: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	foundFriends, err := FindAllFrindsByUserID(allUsersInDB, userID)
	if err != nil {
		handler.log.Printf("cannot find friends: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
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
