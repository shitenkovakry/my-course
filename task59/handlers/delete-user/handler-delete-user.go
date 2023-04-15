package deleteuser

import (
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
