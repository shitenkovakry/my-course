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
