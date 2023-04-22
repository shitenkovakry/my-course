package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pkg/errors"
)

const (
	pathToJSON = "./orders.json"
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
	log log.Logger
}

func readFromJSON() (Orders, error) {
	data, err := os.ReadFile(pathToJSON)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read file")
	}

	var orders Orders

	if err := json.Unmarshal(data, &orders); err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal file")
	}

	return orders, nil
}

func writeToJSON(orders Orders) error {
	data, err := json.Marshal(orders)
	if err != nil {
		return err
	}

	if err := os.WriteFile(pathToJSON, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func NewDispatcher(log log.Logger) *Dispatcher {
	return &Dispatcher{
		log: log,
	}
}

func (dispatcher *Dispatcher) obtainNextID(newOrder *Order) (Orders, int, error) {
	allOrdersInDB, err := dispatcher.ReadOrders()
	if err != nil {
		return nil, 0, errors.Wrap(err, "cannot read all orders")
	}

	nextID := 0

	for index := 0; index < len(allOrdersInDB); index++ {
		order := allOrdersInDB[index]
		if newOrder.ID == order.ID {
			return allOrdersInDB, nextID, errors.Wrapf(ErrAlready, "in db id = %s", newOrder.ID)
		}

		if nextID < order.ID {
			nextID = order.ID
		}
	}

	return allOrdersInDB, nextID + 1, nil
}

func (dispatcher *Dispatcher) CreateOrder(newOrder *Order) (*Order, error) {
	allOrdersInDB, nextID, err := dispatcher.obtainNextID(newOrder)
	if err != nil {
		return nil, errors.Wrap(err, "cannot obtain next ID")
	}

	newOrder.ID = nextID
	allOrdersInDB = append(allOrdersInDB, newOrder)

	if err := writeToJSON(allOrdersInDB); err != nil {
		return nil, errors.Wrap(err, "cannot write into file")
	}

	return newOrder, nil
}

func (dispatcher *Dispatcher) findOrderByID(orderID int) (Orders, *Order, error) {
	allOrdersInDB, err := dispatcher.ReadOrders()
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot read")
	}

	for index := 0; index < len(allOrdersInDB); index++ {
		order := allOrdersInDB[index]

		if orderID == order.ID {
			return allOrdersInDB, order, nil
		}
	}

	return nil, nil, errors.Wrapf(ErrNotFound, "can not find this order = %s")
}

func (dispatcher *Dispatcher) UpdateAddress(orderID int, address string) (*Order, error) {
	allOrdersInDB, order, err := dispatcher.findOrderByID(orderID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot find order by ID")
	}

	order.Address = address

	if err := writeToJSON(allOrdersInDB); err != nil {
		return nil, errors.Wrap(err, "cannot write into file")
	}

	return order, nil
}

func (dispatcher *Dispatcher) UpdateTelephone(orderID int, telephone string) (*Order, error) {
	allOrdersInDB, order, err := dispatcher.findOrderByID(orderID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot find order by ID")
	}

	order.Telephone = telephone

	if err := writeToJSON(allOrdersInDB); err != nil {
		return nil, errors.Wrap(err, "cannot write into file")
	}

	return order, nil
}

func (dispatcher *Dispatcher) ReadOrders() (Orders, error) {
	return readFromJSON()
}

func (dispatcher *Dispatcher) UpdateStatus(orderID int, status string) (*Order, error) {
	allOrdersInDB, order, err := dispatcher.findOrderByID(orderID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot find order by ID")
	}

	order.Status = status

	if err := writeToJSON(allOrdersInDB); err != nil {
		return nil, errors.Wrap(err, "cannot write into file")
	}

	return order, nil
}
