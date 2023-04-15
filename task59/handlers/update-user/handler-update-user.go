package updateuser

import (
	"curse/task59/logger"
	"curse/task59/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type UpdateUser struct {
	Age int `json:"new age"`
}

type UserActionsForHandlerUpdateUser interface {
	UpdateAgeOfUser(userID int, age int) (*models.User, error)
}

type HandlerForUpdateUser struct {
	log         logger.Logger
	userActions UserActionsForHandlerUpdateUser
}

func NewHandlerForUpdateUserByAge(log logger.Logger, userActions UserActionsForHandlerUpdateUser) *HandlerForUpdateUser {
	result := &HandlerForUpdateUser{
		log:         log,
		userActions: userActions,
	}

	return result
}

func (handler *HandlerForUpdateUser) prepareRequest(request *http.Request) (*models.User, error) {
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

	var updateUserFromClient *UpdateUser

	if err := json.Unmarshal(body, &updateUserFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	updatedUser := &models.User{
		ID:  userID,
		Age: updateUserFromClient.Age,
	}

	return updatedUser, nil
}

func (handler *HandlerForUpdateUser) sendResponse(write http.ResponseWriter, updatedUser *models.User) {
	updatedUserMarshaled, err := json.Marshal(updatedUser)
	if err != nil {
		handler.log.Printf("can not marshal created user: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := write.Write(updatedUserMarshaled); err != nil {
		handler.log.Printf("can not send to client updated user: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerForUpdateUser) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	shouldUpdateUser, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		write.WriteHeader(http.StatusBadRequest)

		return
	}

	updatedUser, err := handler.userActions.UpdateAgeOfUser(shouldUpdateUser.ID, shouldUpdateUser.Age)
	if err != nil {
		handler.log.Printf("can not update user: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(write, updatedUser)

}
