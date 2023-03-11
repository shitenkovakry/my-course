package handlers

import (
	"curse/practice30/datasource"
	"encoding/json"
	"io"
	"net/http"
)

type HandlerDeleteUser struct {
	log Logger
}

func NewHandlerDeleteUser(log Logger) *HandlerDeleteUser {
	return &HandlerDeleteUser{
		log: log,
	}
}

func (handler *HandlerDeleteUser) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		deleteUser *DeleteUserInfo
	)

	allUsersInDB, err := datasource.ReadAllUsers()
	if err != nil {
		handler.log.Printf("cannot unmarshal: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.Unmarshal(body, &deleteUser); err != nil {
		handler.log.Printf("cannot unmarshal body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	allFriendsWithoutDeletedUser, err := datasource.UserDeletionResult(allUsersInDB, deleteUser.ID)
	if err != nil {
		handler.log.Printf("cannot delete user: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := datasource.SaveResultIntoFile(allFriendsWithoutDeletedUser); err != nil {
		handler.log.Printf("cannot save: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(deleteUser.ID)
	if err != nil {
		handler.log.Printf("cannot marshal userDeleted")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write([]byte(response)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}
