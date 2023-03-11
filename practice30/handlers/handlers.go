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
	ErrAlready  = errors.New("already")
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

func UpdateUsersInDBWithGivenUserByAge(allUsersInDB Users, userID, age int) (Users, error) {
	userWhoseAgeShouldBeUpdate, err := FindUserByID(allUsersInDB, userID)
	if err != nil {
		return nil, err
	}

	updatedUsers := UpdateAgeOfUser(allUsersInDB, userWhoseAgeShouldBeUpdate.ID, age)

	return updatedUsers, nil
}

func MakeFriendsForSourceUserAndTargetUser(allUsersInDB Users, sourceID int, tergetID int) (*User, *User, error) {
	sourceUser, err := FindUserByID(allUsersInDB, sourceID)
	if err != nil {
		return nil, nil, err
	}

	targetUser, err := FindUserByID(allUsersInDB, tergetID)
	if err != nil {
		return nil, nil, err
	}

	sourceUser.Friends = append(sourceUser.Friends, targetUser.ID)
	targetUser.Friends = append(targetUser.Friends, sourceUser.ID)

	return sourceUser, targetUser, nil
}

func UserDeletionResult(allUsersInDB Users, idOfDeletUser int) (Users, error) {
	userDeleted, err := FindUserByID(allUsersInDB, idOfDeletUser)
	if err != nil {
		return nil, err
	}

	allUsersInDBWithoutDeletedUser := DeleteUser(allUsersInDB, userDeleted.ID)
	allFriendsWithoutDeletedUser := DeleteUserFromFriends(allUsersInDBWithoutDeletedUser, userDeleted.ID)

	return allFriendsWithoutDeletedUser, nil
}
