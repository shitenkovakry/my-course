package handlers

import (
	"curse/practice30/datasource"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

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

	if err := json.Unmarshal(body, &infoOfUpdateUser); err != nil {
		handler.log.Printf("cannot unmarshal body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	updatedUsers, err := datasource.UpdateUsersInDBWithGivenUserByAge(allUsersInDB, userID, infoOfUpdateUser.Age)
	if err != nil {
		handler.log.Printf("cannot update: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := datasource.SaveResultIntoFile(updatedUsers); err != nil {
		handler.log.Printf("cannot save: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	response := fmt.Sprintln("пользователь успешно обновлен")

	if _, err := writer.Write([]byte(response)); err != nil {
		handler.log.Printf("err = %v", err)
	}

}
