package dispatcher

import (
	"encoding/json"
	"os"

	"curse/task60.3/logger"
	"curse/task60.3/models"

	"github.com/pkg/errors"
)

const (
	pathToJSON = "./orders.json"
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

func NewDispatcher(log logger.Logger, db DB) *Dispatcher {
	return &Dispatcher{
		log: log,
		db:  db,
	}
}

// func (dispatcher *Dispatcher) obtainNextID(newOrder *models.Order) (models.Orders, int, error) {
// 	allOrdersInDB, err := dispatcher.ReadOrders()
// 	if err != nil {
// 		return nil, 0, errors.Wrap(err, "cannot read all orders")
// 	}

// 	nextID := 0

// 	for index := 0; index < len(allOrdersInDB); index++ {
// 		order := allOrdersInDB[index]
// 		if newOrder.ID == order.ID {
// 			return allOrdersInDB, nextID, errors.Wrapf(ErrAlready, "in db id = %s", newOrder.ID)
// 		}

// 		if nextID < order.ID {
// 			nextID = order.ID
// 		}
// 	}

// 	return allOrdersInDB, nextID + 1, nil
// }

func (dispatcher *Dispatcher) Create(newOrder *models.Order) (*models.Order, error) {
	// allOrdersInDB, nextID, err := dispatcher.obtainNextID(newOrder)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "cannot obtain next ID")
	// }

	// newOrder.ID = nextID
	// allOrdersInDB = append(allOrdersInDB, newOrder)

	// if err := writeToJSON(allOrdersInDB); err != nil {
	// 	return nil, errors.Wrap(err, "cannot write into file")
	// }

	insertedOrder, err := dispatcher.db.Insert(newOrder)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create order")
	}

	return insertedOrder, nil
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
	// allOrdersInDB, order, err := dispatcher.findOrderByID(orderID)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "cannot find order by ID")
	// }

	// order.Address = address

	// if err := writeToJSON(allOrdersInDB); err != nil {
	// 	return nil, errors.Wrap(err, "cannot write into file")
	// }

	updatedOrder, err := dispatcher.db.UpdateAddress(orderID, address)
	if err != nil {
		return nil, errors.Wrap(err, "can not update order")
	}

	return updatedOrder, nil
}

func (dispatcher *Dispatcher) UpdateTelephone(orderID int, telephone string) (*models.Order, error) {
	/*allOrdersInDB, order, err := dispatcher.findOrderByID(orderID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot find order by ID")
	}

	order.Telephone = telephone

	if err := writeToJSON(allOrdersInDB); err != nil {
		return nil, errors.Wrap(err, "cannot write into file")
	} */

	updatedTelephone, err := dispatcher.db.UpdateTelephone(orderID, telephone)
	if err != nil {
		return nil, errors.Wrapf(err, "can not update telephone")
	}

	return updatedTelephone, nil
}

func (dispatcher *Dispatcher) ReadOrders() (models.Orders, error) {
	//return readFromJSON()

	read, err := dispatcher.db.ReadAll()
	if err != nil {
		return nil, errors.Wrapf(err, "can not read")
	}

	return read, nil
}

func (dispatcher *Dispatcher) UpdateStatus(orderID int, status string) (*models.Order, error) {
	/*allOrdersInDB, order, err := dispatcher.findOrderByID(orderID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot find order by ID")
	}

	order.Status = status

	if err := writeToJSON(allOrdersInDB); err != nil {
		return nil, errors.Wrap(err, "cannot write into file")
	}

	return order, nil */

	updatedStatus, err := dispatcher.db.UpdateStatus(orderID, status)
	if err != nil {
		return nil, errors.Wrapf(err, "can not update status")
	}

	return updatedStatus, nil
}
