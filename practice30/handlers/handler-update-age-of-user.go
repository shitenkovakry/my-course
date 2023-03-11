package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

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

func NewHandlerUpdateUserAge(log Logger) *HandlerUpdateUserAge {
	return &HandlerUpdateUserAge{
		log: log,
	}
}

func (handler *HandlerUpdateUserAge) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		infoOfUpdateUser *UpdateAgeInfo
	)

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

	updatedUsers, err := UpdateUsersInDBWithGivenUserByAge(allUsersInDB, userID, infoOfUpdateUser.Age)
	if err != nil {
		handler.log.Printf("cannot update: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := SaveResultIntoFile(updatedUsers); err != nil {
		handler.log.Printf("cannot save: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	response := fmt.Sprintln("пользователь успешно обновлен")

	if _, err := writer.Write([]byte(response)); err != nil {
		handler.log.Printf("err = %v", err)
	}

}
