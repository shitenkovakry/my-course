package deleteuser

import (
	"curse/task59/logger"
	"curse/task59/models"
)

type DeleteUser struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type UserActionsForHandlerDeleteUser interface {
	Delete(userID int) (*models.Users, error)
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
