package updateuser

import (
	"curse/practice31/logger"
	"curse/practice31/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type UpdateUserByAge struct {
	Age int `json:"new age"`
}

type UserActionsForHandlerUpdateUserByAge interface {
	UpdateAge(userID int, age int) (*models.User, error)
}

type HandlerForUpdateAge struct {
	log         logger.Logger
	userActions UserActionsForHandlerUpdateUserByAge
}

func NewHandlerForUpdateUserByAge(log logger.Logger, userActions UserActionsForHandlerUpdateUserByAge) *HandlerForUpdateAge {
	result := &HandlerForUpdateAge{
		log:         log,
		userActions: userActions,
	}

	return result
}

func (handler *HandlerForUpdateAge) prepareRequest(request *http.Request) (*models.User, error) {
	userIDParam := chi.URLParam(request, "user_id")
	userID, err := strconv.Atoi(userIDParam)

	if err != nil {
		handler.log.Printf("err = %v", err)

		return nil, err
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body: %v", err)

		return nil, err
	}

	var newAgeFromClient *UpdateUserByAge

	if err := json.Unmarshal(body, &newAgeFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newAge := &models.User{
		ID:  userID,
		Age: newAgeFromClient.Age,
	}

	return newAge, nil
}

func (handler *HandlerForUpdateAge) sendResponse(writer http.ResponseWriter, updatedAge *models.User) {
	updatedAgeMarshaled, err := json.Marshal(updatedAge)
	if err != nil {
		handler.log.Printf("cannot marshal updated age: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(updatedAgeMarshaled); err != nil {
		handler.log.Printf("cannot send to client updated age: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerForUpdateAge) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldUpdateAge, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("cannot prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	updatedAge, err := handler.userActions.UpdateAge(shouldUpdateAge.ID, shouldUpdateAge.Age)
	if err != nil {
		handler.log.Printf("cannot update age: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, updatedAge)
}
