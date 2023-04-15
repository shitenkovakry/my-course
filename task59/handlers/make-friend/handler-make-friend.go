package makefriend

import (
	"curse/task59/logger"
	"curse/task59/models"
)

type MakeFriend struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}

type UserActionsForHandlerMakeFriendForUsers interface {
	MakeFriend(sourceID, targetID int) (*models.User, *models.User, error)
}

type HandlerForMakeFriendForUsers struct {
	log         logger.Logger
	userActions UserActionsForHandlerMakeFriendForUsers
}

func NewHandlerForMakeFriend(log logger.Logger, userActions UserActionsForHandlerMakeFriendForUsers) *HandlerForMakeFriendForUsers {
	result := &HandlerForMakeFriendForUsers{
		log:         log,
		userActions: userActions,
	}

	return result
}
