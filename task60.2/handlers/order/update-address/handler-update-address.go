package updateaddress

import (
	"curse/task60.2/logger"
	"curse/task60.2/models"
	"encoding/json"
	"io"
	"net/http"
)

type UpdateAddress struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
}

type OrderActionsForHandlerUpdateAddress interface {
	UpdateAddress(orderID int, address string) (*models.Order, error)
}

type HandlerForUpdateAddress struct {
	log          logger.Logger
	orderActions OrderActionsForHandlerUpdateAddress
}

func NewHandlerForUpdateAddress(log logger.Logger, orderActions OrderActionsForHandlerUpdateAddress) *HandlerForUpdateAddress {
	result := &HandlerForUpdateAddress{
		log:          log,
		orderActions: orderActions,
	}

	return result
}

func (handler *HandlerForUpdateAddress) prepareRequest(request *http.Request) (*models.Order, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body: %v", err)

		return nil, err
	}

	var newAddressFromClient *UpdateAddress

	if err := json.Unmarshal(body, &newAddressFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newAddress := &models.Order{
		ID:      newAddressFromClient.ID,
		Address: newAddressFromClient.Address,
	}

	return newAddress, nil
}

func (handler *HandlerForUpdateAddress) sendResponse(writer http.ResponseWriter, updatedAddress *models.Order) {
	updatedAddressMarshaled, err := json.Marshal(updatedAddress)
	if err != nil {
		handler.log.Printf("cannot marshal updated address: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(updatedAddressMarshaled); err != nil {
		handler.log.Printf("cannot send to client updated address: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerForUpdateAddress) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldUpdateAddress, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("cannot prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	updatedAddress, err := handler.orderActions.UpdateAddress(shouldUpdateAddress.ID, shouldUpdateAddress.Address)
	if err != nil {
		handler.log.Printf("cannot update address: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, updatedAddress)
}
