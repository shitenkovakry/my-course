package main

import (
	"errors"
)

type User struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type Users []*User

type MakeNewFriendInfo struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}

type DeleteUserInfo struct {
	ID int `json:"target_id"`
}

type UpdateAgeInfo struct {
	Age int `json:"new_age"`
}

var (
	ErrNotFound = errors.New("not found")
)

func FindUserByID(allUsersInDB Users, id int) (*User, error) {
	for i := 0; i <= len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID == id {
			return user, nil
		}
	}

	return nil, ErrNotFound
}
