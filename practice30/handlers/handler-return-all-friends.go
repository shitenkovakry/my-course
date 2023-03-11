package handlers

import (
	"curse/practice30/datasource"
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

func (handler *HandlerReturnAllFriendsID) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userIDParam := chi.URLParam(request, "user_id")
	userID, err := strconv.Atoi(userIDParam)

	if err != nil {
		handler.log.Printf("err = %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	allUsersInDB, err := datasource.ReadAllUsers()
	if err != nil {
		handler.log.Printf("cannot unmarshal: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	foundFriends, err := datasource.FindAllFrindsByUserID(allUsersInDB, userID)
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
