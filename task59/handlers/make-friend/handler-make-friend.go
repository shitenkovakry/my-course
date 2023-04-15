package makefriend

import "curse/task59/models"

type MakeFriend struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}

type UserActionsForHandlerMakeFriendForUsers interface {
	MakeFriend(sourceID, targetID int) (*models.User, *models.User, error)
}
