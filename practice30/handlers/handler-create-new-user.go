package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
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

func CreateNewUser(allUsersInDB Users, newUser *User) (Users, error) {
	nextID := 0
	for _, user := range allUsersInDB {
		if newUser.Email == user.Email {
			return nil, errors.Wrapf(ErrAlready, "in db email = %s", newUser.Email)
		}

		if nextID < user.ID {
			nextID = user.ID
		}
	}

	newUser.ID = nextID + 1

	allUsersInDB = append(allUsersInDB, newUser)

	return allUsersInDB, nil
}

func (handler *HandlerCreateUser) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		newUser *User
	)

	allUsersInDB, err := ReadAllUsers()
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

	usersInDBWithNewUser, err := CreateNewUser(allUsersInDB, newUser)
	if err != nil {
		handler.log.Printf("cannot create new user:, %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := SaveResultIntoFile(usersInDBWithNewUser); err != nil {
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
