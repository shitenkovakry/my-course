package main

import (
	"encoding/json"
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
	/*
		if _, err := writer.Write([]byte(`{"email":"my@mail.ru","name":"some name","age":"24","friends":[]}`)); err != nil {
			handler.log.Printf("err = %v", err)
		}
	*/

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

		if nextID < newUser.ID {
			nextID = newUser.ID
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

	response, err := json.Marshal(newUser)
	if err != nil {
		handler.log.Printf("cannot marshal newUser")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(response); err != nil {
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
