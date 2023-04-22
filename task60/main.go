package main

import (
	"log"

	"github.com/pkg/errors"
)

var (
	ErrAlready  = errors.New("already")
	ErrNotFound = errors.New("not found")
)

type Order struct {
	ID        int
	Address   string
	Telephone string
	Status    string
}

type Orders []*Order

type Dispatcher struct {
	orders Orders
	log    log.Logger
}

func NewDispatcher(log log.Logger) *Dispatcher {
	return &Dispatcher{
		log:    log,
		orders: make(Orders, 0),
	}
}

func (dispatcher *Dispatcher) obtainNextID(newOrder *Order) (int, error) {
	allOrdersInDB := dispatcher.orders
	nextID := 0

	for index := 0; index < len(allOrdersInDB); index++ {
		order := allOrdersInDB[index]
		if newOrder.ID == order.ID {
			return nextID, errors.Wrapf(ErrAlready, "in db id = %s", newOrder.ID)
		}

		if nextID < order.ID {
			nextID = order.ID
		}
	}

	return nextID + 1, nil
}

func (dispatcher *Dispatcher) CreateOrder(newOrder *Order) (*Order, error) {
	nextID, err := dispatcher.obtainNextID(newOrder)
	if err != nil {
		return nil, errors.Wrap(err, "cannot obtain next ID")
	}

	newOrder.ID = nextID
	dispatcher.orders = append(dispatcher.orders, newOrder)

	return newOrder, nil
}

func (dispatcher *Dispatcher) UpdateAddress(orderID int, address string) (*Order, error) {
	allOrdersInDB := dispatcher.orders

	for index := 0; index < len(allOrdersInDB); index++ {
		order := allOrdersInDB[index]

		if orderID == order.ID {
			order.Address = address

			return order, nil
		}
	}

	return nil, errors.Wrapf(ErrNotFound, "can not find this order = %s")
}

func (dispatcher *Dispatcher) UpdateTelephone(orderID int, telephone string) (*Order, error) {
	allOrdersInDB := dispatcher.orders

	for index := 0; index < len(allOrdersInDB); index++ {
		order := allOrdersInDB[index]

		if orderID == order.ID {
			order.Telephone = telephone

			return order, nil
		}
	}

	return nil, errors.Wrapf(ErrNotFound, "can not find this order = %s")
}

func (dispatcher *Dispatcher) ReadOrders() (Orders, error) {
	return dispatcher.orders, nil
}

func (dispatcher *Dispatcher) UpdateStatus(orderID int, status string) (*Order, error) {
	allOrdersInDB := dispatcher.orders

	for index := 0; index < len(allOrdersInDB); index++ {
		order := allOrdersInDB[index]

		if orderID == order.ID {
			order.Status = status

			return order, nil
		}
	}

	return nil, errors.Wrapf(ErrNotFound, "can not find this status = %s")
}
