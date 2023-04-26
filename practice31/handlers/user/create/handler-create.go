package user

import (
	"curse/practice31/logger"
	"curse/practice31/models"
	"encoding/json"
	"io"
	"net/http"
)

type NewUser struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type UserActionsForHandlerCreateUser interface {
	Create(newUser *models.User) (*models.User, error)
}

type HandlerForCreateUser struct {
	log         logger.Logger
	userActions UserActionsForHandlerCreateUser
}

func NewHandlerForCreateUser(log logger.Logger, userActions UserActionsForHandlerCreateUser) *HandlerForCreateUser {
	result := &HandlerForCreateUser{
		log:         log,
		userActions: userActions,
	}

	return result
}

func (handler *HandlerForCreateUser) prepareRequest(request *http.Request) (*models.User, error) {
	defer func() {
		if err := request.Body.Close(); err != nil {
			handler.log.Printf("cannot close body: %v", err)
		}
	}()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body: %v", err)

		return nil, err
	}

	var newUserFromClient *NewUser

	if err := json.Unmarshal(body, &newUserFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newUser := &models.User{
		Email:   newUserFromClient.Email,
		Name:    newUserFromClient.Name,
		Age:     newUserFromClient.Age,
		Friends: newUserFromClient.Friends,
	}

	return newUser, nil
}

func (handler *HandlerForCreateUser) sendResponse(write http.ResponseWriter, createdUser *models.User) {
	createdUserMarshaled, err := json.Marshal(createdUser)
	if err != nil {
		handler.log.Printf("cannot marshal created user: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := write.Write(createdUserMarshaled); err != nil {
		handler.log.Printf("cannot send to client created user: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerForCreateUser) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	newUser, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("cannot prepare request: %v", err)
		write.WriteHeader(http.StatusBadRequest)

		return
	}

	createdUser, err := handler.userActions.Create(newUser)
	if err != nil {
		handler.log.Printf("cannot create user: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(write, createdUser)
}
