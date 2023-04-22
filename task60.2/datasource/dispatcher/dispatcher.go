package dispatcher

import (
	"encoding/json"
	"os"

	"curse/task60.2/logger"
	"curse/task60.2/models"

	"github.com/pkg/errors"
)

const (
	pathToJSON = "./orders.json"
)

var (
	ErrAlready  = errors.New("already")
	ErrNotFound = errors.New("not found")
)

type Dispatcher struct {
	log logger.Logger
}

func readFromJSON() (models.Orders, error) {
	data, err := os.ReadFile(pathToJSON)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read file")
	}

	var orders models.Orders

	if err := json.Unmarshal(data, &orders); err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal file")
	}

	return orders, nil
}

func writeToJSON(orders models.Orders) error {
	data, err := json.Marshal(orders)
	if err != nil {
		return err
	}

	if err := os.WriteFile(pathToJSON, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func NewDispatcher(log logger.Logger) *Dispatcher {
	return &Dispatcher{
		log: log,
	}
}

func (dispatcher *Dispatcher) obtainNextID(newOrder *models.Order) (models.Orders, int, error) {
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

func (dispatcher *Dispatcher) Create(newOrder *models.Order) (*models.Order, error) {
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

func (dispatcher *Dispatcher) findOrderByID(orderID int) (models.Orders, *models.Order, error) {
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

func (dispatcher *Dispatcher) UpdateAddress(orderID int, address string) (*models.Order, error) {
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

func (dispatcher *Dispatcher) UpdateTelephone(orderID int, telephone string) (*models.Order, error) {
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

func (dispatcher *Dispatcher) ReadOrders() (models.Orders, error) {
	return readFromJSON()
}

func (dispatcher *Dispatcher) UpdateStatus(orderID int, status string) (*models.Order, error) {
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
