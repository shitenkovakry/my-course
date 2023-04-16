package users

import (
	"curse/task59/models"
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrAlready  = errors.New("already")
)

type Registration struct {
	allUsersInDB models.Users
}

func (registration *Registration) Create(newUser *models.User) (*models.User, error) {
	nextID := 0
	for _, user := range registration.allUsersInDB {
		if newUser.Email == user.Email {
			return nil, errors.Wrapf(ErrAlready, "in db email = %s", newUser.Email)
		}

		if nextID < user.ID {
			nextID = user.ID
		}
	}

	newUser.ID = nextID + 1

	registration.allUsersInDB = append(registration.allUsersInDB, newUser)

	return newUser, nil
}

func (registration *Registration) UpdateAgeOfUser(userID int, age int) (*models.User, error) {
	allUsersInDB := registration.allUsersInDB

	for i := 0; i < len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID == userID {
			user.Age = age

			return user, nil
		}
	}

	return nil, errors.Wrapf(ErrNotFound, "can not find this person = %s")
}

func (registration *Registration) Delete(userID int) (*models.User, error) {
	allUsersInDB := registration.allUsersInDB
	newArray := models.Users{}

	var deletedUser *models.User

	//for _, user := range allUsersInDB { - for dumbs
	for i := 0; i < len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID == userID {
			deletedUser = user

			continue
		}

		newArray = append(newArray, user)
	}

	registration.allUsersInDB = newArray

	if deletedUser != nil {
		return deletedUser, nil
	}

	return nil, errors.Wrapf(ErrNotFound, "can not find this person = %s")
}

func (registretion *Registration) FindUserByID(id int) (*models.User, error) {
	allUsersInDB := registretion.allUsersInDB

	for index := 0; index <= len(allUsersInDB); index++ {
		user := allUsersInDB[index]

		if user.ID == id {
			return user, nil
		}
	}

	return nil, ErrNotFound
}

func (registration *Registration) MakeFriend(sourceID, targetID int) (string, error) {
	sourceUser, err := registration.FindUserByID(sourceID)
	if err != nil {
		return "", err
	}

	targetUser, err := registration.FindUserByID(targetID)
	if err != nil {
		return "", err
	}

	sourceUser.Friends = append(sourceUser.Friends, targetUser.ID)
	targetUser.Friends = append(targetUser.Friends, sourceUser.ID)

	return fmt.Sprintf("%s and %s are friends", sourceUser.Name, targetUser.Name), nil
}

func New() *Registration {
	return &Registration{
		allUsersInDB: models.Users{},
	}
}
