package datasource

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type User struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type Users []*User

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

func CreateNewUser(allUsersInDB Users, newUser *User) (Users, error) {
	nextID := 0
	for _, user := range allUsersInDB {
		if newUser.Email == user.Email {
			return nil, errors.Wrapf(ErrAlready, "in db email = %s", newUser.Email)
		}

		if nextID < user.ID {
			nextID = user.ID
		}
	}

	newUser.ID = nextID + 1

	allUsersInDB = append(allUsersInDB, newUser)

	return allUsersInDB, nil
}

func DeleteUser(allUsersInDB Users, id int) Users {
	newArray := Users{}

	//for _, user := range allUsersInDB { - for dumbs
	for i := 0; i < len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID != id {
			newArray = append(newArray, user)
		} else if user.ID == id {
			continue
		}
	}

	return newArray
}

func DeleteUserFromFriends(allUsersInDB Users, id int) Users {
	newArrayOfFriends := Users{}

	for i := 0; i < len(allUsersInDB); i++ {
		user := allUsersInDB[i]
		arrayOfFriendsOfUser := []int{}

		for j := 0; j < len(user.Friends); j++ {
			friendID := user.Friends[j]

			if friendID != id {
				arrayOfFriendsOfUser = append(arrayOfFriendsOfUser, friendID)
			} else if friendID == id {
				continue
			}
		}

		user.Friends = arrayOfFriendsOfUser
		newArrayOfFriends = append(newArrayOfFriends, user)
	}

	return newArrayOfFriends
}

func FindAllFrindsByUserID(allUsersInDB Users, userID int) (Users, error) {
	var foundUser *User

	for _, user := range allUsersInDB {
		if user.ID == userID {
			foundUser = user

			break
		}
	}

	if foundUser == nil {
		return nil, ErrNotFound
	}

	foundFriends := Users{}

	for _, userID := range foundUser.Friends {
		for _, user := range allUsersInDB {
			if user.ID == userID {
				foundFriends = append(foundFriends, user)

				break
			}
		}
	}

	return foundFriends, nil
}

func UpdateAgeOfUser(allUsersInDB Users, id int, age int) Users {
	newArray := Users{}
	for i := 0; i < len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID == id {
			user.Age = age
		}

		newArray = append(newArray, user)
	}

	return newArray
}
