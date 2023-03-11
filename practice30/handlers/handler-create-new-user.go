package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
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
		allUsersInDB Users
		newUser      *User
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

	response, err := json.Marshal(newUser.ID)
	if err != nil {
		handler.log.Printf("cannot marshal newUser")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(response); err != nil {
		handler.log.Printf("err = %v", err)
	}

	writer.WriteHeader(http.StatusCreated)
}
