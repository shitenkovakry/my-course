package create

import (
	"curse/task59/logger"
	"curse/task60.2/models"
	"encoding/json"
	"io"
	"net/http"
)

type NewOrder struct {
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
}

type OrderActionsForHandlerCreateOrder interface {
	Create(newOrder *models.Order) (*models.Order, error)
}

type HandlerForCreateOrder struct {
	log          logger.Logger
	orderActions OrderActionsForHandlerCreateOrder
}

func NewHandlerForCreateOrder(log logger.Logger, orderActions OrderActionsForHandlerCreateOrder) *HandlerForCreateOrder {
	result := &HandlerForCreateOrder{
		log:          log,
		orderActions: orderActions,
	}

	return result
}

func (handler *HandlerForCreateOrder) prepareRequest(request *http.Request) (*models.Order, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body: %v", err)

		return nil, err
	}

	var newOrderFromClient *NewOrder

	if err := json.Unmarshal(body, &newOrderFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newOrder := &models.Order{
		Address:   newOrderFromClient.Address,
		Telephone: newOrderFromClient.Telephone,
	}

	return newOrder, nil
}

func (handler *HandlerForCreateOrder) sendResponse(write http.ResponseWriter, createdOrder *models.Order) {
	createdOrderMarshaled, err := json.Marshal(createdOrder)
	if err != nil {
		handler.log.Printf("cannot marshal created order: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := write.Write(createdOrderMarshaled); err != nil {
		handler.log.Printf("cannot send to client created order: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerForCreateOrder) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	newOrder, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("cannot prepare request: %v", err)
		write.WriteHeader(http.StatusBadRequest)

		return
	}

	createdOrder, err := handler.orderActions.Create(newOrder)
	if err != nil {
		handler.log.Printf("cannot create order: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(write, createdOrder)
}
