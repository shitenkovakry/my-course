package updatetelephone

import (
	"curse/task60.2/logger"
	"curse/task60.2/models"
	"encoding/json"
	"io"
	"net/http"
)

type UpdateTelephone struct {
	ID        int    `json:"id"`
	Telephone string `json:"telephone"`
}

type OrderActionsForHandlerUpdateTelephone interface {
	UpdateTelephone(orderID int, telephone string) (*models.Order, error)
}

type HandlerForUpdateTelephone struct {
	log          logger.Logger
	orderActions OrderActionsForHandlerUpdateTelephone
}

func NewHandlerForUpdateTelephone(log logger.Logger, orderActions OrderActionsForHandlerUpdateTelephone) *HandlerForUpdateTelephone {
	result := &HandlerForUpdateTelephone{
		log:          log,
		orderActions: orderActions,
	}

	return result
}

func (handler *HandlerForUpdateTelephone) prepareRequest(request *http.Request) (*models.Order, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body: %v", err)

		return nil, err
	}

	var newTelephoneFromClient *UpdateTelephone

	if err := json.Unmarshal(body, &newTelephoneFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newTelephone := &models.Order{
		ID:        newTelephoneFromClient.ID,
		Telephone: newTelephoneFromClient.Telephone,
	}

	return newTelephone, nil
}

func (handler *HandlerForUpdateTelephone) sendResponse(writer http.ResponseWriter, updatedTelephone *models.Order) {
	updatedTelephoneMarshaled, err := json.Marshal(updatedTelephone)
	if err != nil {
		handler.log.Printf("cannot marshal updated telephone: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(updatedTelephoneMarshaled); err != nil {
		handler.log.Printf("cannot send to client updated telephone: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerForUpdateTelephone) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldUpdateTelephone, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("cannot prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	updatedTelephone, err := handler.orderActions.UpdateTelephone(shouldUpdateTelephone.ID, shouldUpdateTelephone.Telephone)
	if err != nil {
		handler.log.Printf("cannot update telephone: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, updatedTelephone)
}
