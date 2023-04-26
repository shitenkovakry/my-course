package users

import (
	"curse/practice31/logger"
	"curse/practice31/models"

	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrAlready  = errors.New("already")
)

type DB interface {
	Insert(user *models.User) (*models.User, error)
	UpdateAge(id int, age int) (*models.User, error)
	DeleteUser(id int) (*models.User, error)
}

type Registration struct {
	log logger.Logger
	db  DB
}

func New(log logger.Logger, db DB) *Registration {
	return &Registration{
		log: log,
		db:  db,
	}
}

func (registration *Registration) Create(newUser *models.User) (*models.User, error) {
	insertedUser, err := registration.db.Insert(newUser)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create user")
	}

	return insertedUser, nil
}

func (registration *Registration) UpdateAge(userID int, age int) (*models.User, error) {
	updatedAge, err := registration.db.UpdateAge(userID, age)
	if err != nil {
		return nil, errors.Wrap(err, "can not update user")
	}

	return updatedAge, nil
}

func (registration *Registration) Delete(userID int) (*models.User, error) {
	deletedUser, err := registration.db.DeleteUser(userID)
	if err != nil {
		return nil, errors.Wrap(err, "can not delete user")
	}

	return deletedUser, nil
}
