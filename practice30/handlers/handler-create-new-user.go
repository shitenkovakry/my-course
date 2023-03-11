package handlers

import (
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

	nextID := 0
	for _, user := range allUsersInDB {
		if newUser.Email == user.Email {
			handler.log.Printf("user have created = %d", newUser.ID)
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		if nextID < user.ID {
			nextID = user.ID
		}
	}

	newUser.ID = nextID + 1

	allUsersInDB = append(allUsersInDB, newUser)
	if err := SaveResultIntoFile(allUsersInDB); err != nil {
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
