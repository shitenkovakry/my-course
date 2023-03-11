package handlers

import (
	"encoding/json"
	"errors"
	"os"
)

type Logger interface {
	Printf(format string, v ...any)
}

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

func ReadAllUsers() (Users, error) {
	var allUsersInDB Users

	dataReadAllUsersInDB, err := os.ReadFile("./database.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dataReadAllUsersInDB, &allUsersInDB); err != nil {
		return nil, err
	}

	return allUsersInDB, nil
}

func SaveResultIntoFile(result Users) error {
	dataWriteAllUsersInDB, err := json.Marshal(result)
	if err != nil {
		return err
	}

	if err := os.WriteFile("./database.json", dataWriteAllUsersInDB, os.ModePerm); err != nil {
		return err
	}

	return nil
}
