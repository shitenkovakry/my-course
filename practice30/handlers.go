package main

import (
	"encoding/json"
	"errors"
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

type MakeNewFriendInfo struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}

type DeleteUserInfo struct {
	ID int `json:"target_id"`
}

var (
	ErrNotFound = errors.New("not found")
)

type HandlerUpdateUserAge struct {
	log Logger
}

func (handler *HandlerUpdateUserAge) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`{"new age":"28"}`)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}

type HandlerDeleteUser struct {
	log Logger
}

func FindUserByID(allUsersInDB Users, id int) (*User, error) {
	for i := 0; i <= len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID == id {
			return user, nil
		}
	}

	return nil, ErrNotFound
}

func DeleteUser(allUsersInDB Users, id int) Users {
	newArray := Users{}

	//for _, user := range allUsersInDB { - for dumbs
	for i := 0; i < len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID != id {
			newArray = append(newArray, user)
		} else if user.ID == id {
			continue
		}
	}

	return newArray
}

func (handler *HandlerDeleteUser) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		deleteUser   *DeleteUserInfo
		allUsersInDB Users
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

	if err := json.Unmarshal(body, &deleteUser); err != nil {
		handler.log.Printf("cannot unmarshal body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	userDeleted, err := FindUserByID(allUsersInDB, deleteUser.ID)
	if err != nil {
		handler.log.Printf("cannot find user %d", deleteUser.ID)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	allUsersInDBWithoutDeletedUser := DeleteUser(allUsersInDB, userDeleted.ID)

	dataWriteAllUsersInDB, err := json.Marshal(allUsersInDBWithoutDeletedUser)
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

	response, err := json.Marshal(userDeleted.ID)
	if err != nil {
		handler.log.Printf("cannot marshal userDeleted")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write([]byte(response)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}
