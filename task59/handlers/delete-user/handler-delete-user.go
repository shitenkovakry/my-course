package deleteuser

import (
	"curse/task59/logger"
	"curse/task59/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type DeleteUser struct {
	ID string `json:""target_id"`
}

type UserActionsForHandlerDeleteUser interface {
	Delete(userID int) (*models.User, error)
}

type HandlerForDeleteUser struct {
	log         logger.Logger
	userActions UserActionsForHandlerDeleteUser
}

func NewHandlerForDeleteUser(log logger.Logger, userActions UserActionsForHandlerDeleteUser) *HandlerForDeleteUser {
	result := &HandlerForDeleteUser{
		log:         log,
		userActions: userActions,
	}

	return result
}

func (handler *HandlerForDeleteUser) prepareRequest(request *http.Request) (*models.User, error) {
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

	var deleteUserFromClient *DeleteUser

	if err := json.Unmarshal(body, &deleteUserFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	deletedUser := &models.User{
		ID: userID,
	}

	return deletedUser, nil
}
