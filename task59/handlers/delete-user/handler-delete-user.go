package deleteuser

import (
	"curse/task59/logger"
	"curse/task59/models"
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
