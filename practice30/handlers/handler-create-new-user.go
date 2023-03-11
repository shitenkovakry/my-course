package handlers

import (
	"curse/practice30/datasource"
	"encoding/json"
	"io"
	"net/http"
)

type HandlerCreateUser struct {
	log Logger
}

func NewHandlerCreateUser(log Logger) *HandlerCreateUser {
	h := &HandlerCreateUser{
		log: log,
	}

	return h
}

func (handler *HandlerCreateUser) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		newUser *datasource.User
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

	if err := json.Unmarshal(body, &newUser); err != nil {
		handler.log.Printf("cannot unmarshal body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	usersInDBWithNewUser, err := datasource.CreateNewUser(allUsersInDB, newUser)
	if err != nil {
		handler.log.Printf("cannot create new user:, %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := datasource.SaveResultIntoFile(usersInDBWithNewUser); err != nil {
		handler.log.Printf("cannot save: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(newUser.ID)
	if err != nil {
		handler.log.Printf("cannot marshal newUser: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(response); err != nil {
		handler.log.Printf("err = %v", err)
	}

	writer.WriteHeader(http.StatusCreated)
}
