package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
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

type UpdateAgeInfo struct {
	Age int `json:"new_age"`
}

var (
	ErrNotFound = errors.New("not found")
)

func FindUserByID(allUsersInDB Users, id int) (*User, error) {
	for i := 0; i <= len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID == id {
			return user, nil
		}
	}

	return nil, ErrNotFound
}

func UpdateAgeOfUser(allUsersInDB Users, id int, age int) Users {
	newArray := Users{}
	for i := 0; i < len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID == id {
			user.Age = age
		}

		newArray = append(newArray, user)
	}

	return newArray
}

type HandlerUpdateUserAge struct {
	log Logger
}

func (handler *HandlerUpdateUserAge) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		infoOfUpdateUser *UpdateAgeInfo
		allUsersInDB     Users
	)

	userIDParam := chi.URLParam(request, "user_id")
	userID, err := strconv.Atoi(userIDParam)

	if err != nil {
		handler.log.Printf("err = %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

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

	if err := json.Unmarshal(body, &infoOfUpdateUser); err != nil {
		handler.log.Printf("cannot unmarshal body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	userWhoseAgeShouldBeUpdate, err := FindUserByID(allUsersInDB, userID)
	if err != nil {
		if err != nil {
			handler.log.Printf("cannot find user %d", userID)
			writer.WriteHeader(http.StatusBadRequest)

			return
		}
	}

	updatedUsers := UpdateAgeOfUser(allUsersInDB, userWhoseAgeShouldBeUpdate.ID, infoOfUpdateUser.Age)

	dataWriteAllUsersInDB, err := json.Marshal(updatedUsers)
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

	response := fmt.Sprintln("пользователь успешно обновлен")

	if _, err := writer.Write([]byte(response)); err != nil {
		handler.log.Printf("err = %v", err)
	}

}
