package dispatcher

import (
	"curse/task60.3/logger"
	"curse/task60.3/models"

	"github.com/pkg/errors"
)

var (
	ErrAlready  = errors.New("already")
	ErrNotFound = errors.New("not found")
)

type DB interface {
	Insert(order *models.Order) (*models.Order, error)
	UpdateAddress(ID int, address string) (*models.Order, error)
	UpdateTelephone(ID int, telephone string) (*models.Order, error)
	UpdateStatus(ID int, status string) (*models.Order, error)
	ReadAll() (models.Orders, error)
}

type Dispatcher struct {
	log logger.Logger
	db  DB
}

func NewDispatcher(log logger.Logger, db DB) *Dispatcher {
	return &Dispatcher{
		log: log,
		db:  db,
	}
}

func (dispatcher *Dispatcher) Create(newOrder *models.Order) (*models.Order, error) {
	insertedOrder, err := dispatcher.db.Insert(newOrder)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create order")
	}

	return insertedOrder, nil
}

func (dispatcher *Dispatcher) UpdateAddress(orderID int, address string) (*models.Order, error) {
	updatedOrder, err := dispatcher.db.UpdateAddress(orderID, address)
	if err != nil {
		return nil, errors.Wrap(err, "can not update order")
	}

	return updatedOrder, nil
}

func (dispatcher *Dispatcher) UpdateTelephone(orderID int, telephone string) (*models.Order, error) {
	updatedTelephone, err := dispatcher.db.UpdateTelephone(orderID, telephone)
	if err != nil {
		return nil, errors.Wrapf(err, "can not update telephone")
	}

	return updatedTelephone, nil
}

func (dispatcher *Dispatcher) ReadOrders() (models.Orders, error) {
	read, err := dispatcher.db.ReadAll()
	if err != nil {
		return nil, errors.Wrapf(err, "can not read")
	}

	return read, nil
}

func (dispatcher *Dispatcher) UpdateStatus(orderID int, status string) (*models.Order, error) {
	updatedStatus, err := dispatcher.db.UpdateStatus(orderID, status)
	if err != nil {
		return nil, errors.Wrapf(err, "can not update status")
	}

	return updatedStatus, nil
}
