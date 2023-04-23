package readall

import (
	"curse/task60.3/logger"
	"curse/task60.3/models"
	"encoding/json"
	"net/http"
)

type OrderActionsForHandlerReadOrder interface {
	ReadOrders() (models.Orders, error)
}

type HandlerForReadOrder struct {
	log          logger.Logger
	orderActions OrderActionsForHandlerReadOrder
}

func NewHandlerForReadOrder(log logger.Logger, orderActions OrderActionsForHandlerReadOrder) *HandlerForReadOrder {
	result := &HandlerForReadOrder{
		log:          log,
		orderActions: orderActions,
	}

	return result
}

func (handler *HandlerForReadOrder) sendResponse(writer http.ResponseWriter, readOrders models.Orders) {
	data, err := json.Marshal(readOrders)
	if err != nil {
		handler.log.Printf("can not marshal orders: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(data); err != nil {
		handler.log.Printf("can not writer orders: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerForReadOrder) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	orders, err := handler.orderActions.ReadOrders()
	if err != nil {
		handler.log.Printf("can not read orders: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, orders)
}
